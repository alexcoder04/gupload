package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/alexcoder04/friendly/v2"
	"github.com/alexcoder04/friendly/v2/ffiles"
	"github.com/gin-gonic/gin"
)

var (
	DeleteTimeout                = 60 * 5
	AutoDeleteMap map[string]int = map[string]int{}
)

func HandlerIndex(c *gin.Context) {
	RetPage(c, []string{})
}

func HandlerUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		RetPage(c, []string{"upload failed"})
		return
	}

	autodelete := form.Value["autodelete"]
	files := form.File["upload"]

	messages := []string{}
	for _, file := range files {
		filename := file.Filename
		filename = regexp.MustCompile(`^[^a-zA-Z0-9_\-.]+$`).ReplaceAllString(filename, "_")

		if err := c.SaveUploadedFile(file, filepath.Join(SharedDir, filename)); err != nil {
			RetPage(c, []string{"Upload failed"})
			continue
		}

		messages = append(messages, fmt.Sprintf("'%s' was uploaded successfully", filename))
		if len(autodelete) > 0 && autodelete[0] == "on" {
			AutoDeleteMap[filename] = DeleteTimeout
		}

		go func() {
			for {
				newMap := map[string]int{}
				for key, val := range AutoDeleteMap {
					if val <= 0 {
						os.Remove(filepath.Join(SharedDir, key))
						continue
					}
					newMap[key] = val - 5
				}
				AutoDeleteMap = newMap
				time.Sleep(5 * time.Second)
			}
		}()
	}

	RetPage(c, messages)
}

func HandlerQRCode(c *gin.Context) {
	c.Data(http.StatusOK, "image/png", QRCodeData)
}

func HandlerDownload(c *gin.Context) {
	if !ffiles.IsDir(filepath.Join(SharedDir, c.Param("path"))) {
		c.File(filepath.Join(SharedDir, c.Param("path")))
	}
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/?view=%s", c.Param("path")))
}

func HandlerZip(c *gin.Context) {
	if c.Param("path") == "" || !ffiles.IsDir(filepath.Join(SharedDir, c.Param("path"))) {
		c.Redirect(http.StatusBadRequest, "/")
	}

	friendly.CompressFolder(filepath.Join(SharedDir, c.Param("path")), filepath.Join(os.TempDir(), c.Param("path")+".zip"))
	c.File(filepath.Join(os.TempDir(), c.Param("path")+".zip"))
}
