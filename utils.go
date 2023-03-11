package main

import (
	"bytes"
	"errors"
	"html/template"
	"image/png"
	"net"
	"os"
	"path"
	"path/filepath"

	"github.com/alexcoder04/friendly/v2/ffiles"
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

var QRCodeData []byte

type SharedFile struct {
	Name      string
	Path      string
	ExpiresIn int
	IsDir     bool
}

func GetLocalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			if ipnet.IP.IsLoopback() {
				continue
			}

			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("not connected to the internet")
}

func QREncode(data string) error {
	qr, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	err = png.Encode(buf, qr.Image(256))
	if err != nil {
		return err
	}

	QRCodeData = buf.Bytes()
	return nil
}

func GetExpireTime(file string) int {
	if val, ok := AutoDeleteMap[file]; ok {
		return val
	}
	return 0
}

func GetSharedFiles(subfolder string) []SharedFile {
	fileInfos, err := os.ReadDir(filepath.Join(SharedDir, subfolder))
	if err != nil {
		return []SharedFile{}
	}

	files := make([]SharedFile, len(fileInfos))
	for i, f := range fileInfos {
		files[i] = SharedFile{
			Name:      f.Name(),
			Path:      path.Join(subfolder, f.Name()),
			ExpiresIn: GetExpireTime(f.Name()),
			IsDir:     ffiles.IsDir(filepath.Join(SharedDir, subfolder, f.Name())),
		}
	}

	return files
}

func PrepareSharedDir() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		SharedDir = os.TempDir()
		return false
	}

	if ffiles.IsDir(filepath.Join(home, "Temp")) {
		SharedDir = filepath.Join(home, "Temp")
		return true
	}

	SharedDir = os.TempDir()
	return false
}

func LoadTemplate() (*template.Template, error) {
	tmpl := template.New("")

	contents, err := TemplateFiles.ReadFile("templates/index.html")
	if err != nil {
		return nil, err
	}

	_, err = tmpl.Parse(string(contents))
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func RetPage(c *gin.Context, messages []string) {
	Template.Execute(c.Writer, gin.H{
		"hostname": Hostname,
		"messages": messages,
		"view":     c.Request.URL.Query().Get("view"),
		"files":    GetSharedFiles(c.Request.URL.Query().Get("view")),
	})
}
