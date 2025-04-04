package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
)

//go:embed templates/index.html
var TemplateFiles embed.FS

//go:embed static
var StaticFiles embed.FS
var StaticFS http.FileSystem

var (
	SharedDir string
)

func init() {
	// find and prepare shared dir
	if !PrepareSharedDir() {
		err := os.MkdirAll(SharedDir, 0700)
		if err != nil {
			panic("failed to create shared dir")
		}
	}

	// generate qr code
	err := QREncode(fmt.Sprintf("http://%s:%s/", Config.IP, Config.Port))
	if err != nil {
		panic("failed to create qr code")
	}

	// strip prefix from static fs
	sub, err := fs.Sub(StaticFiles, "static")
	if err != nil {
		panic("failed to strip 'static' from static files")
	}
	StaticFS = http.FS(sub)
}
