package main

import (
	"embed"
	"fmt"
	"html/template"
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
	Port      string = "1234"
	Hostname  string
	IP        string
	SharedDir string
	Template  *template.Template
)

func init() {
	// find and prepare shared dir
	if !PrepareSharedDir() {
		err := os.MkdirAll(SharedDir, 0700)
		if err != nil {
			panic("failed to create shared dir")
		}
	}

	// get hostname
	var err error
	Hostname, err = os.Hostname()
	if err != nil {
		Hostname = "unknown-pc"
	}

	// get ip
	IP, err = GetLocalIP()
	if err != nil {
		panic("failed to get local ip address")
	}

	// generate qr code
	err = QREncode(fmt.Sprintf("http://%s:%s/", IP, Port))
	if err != nil {
		panic("failed to create qr code")
	}

	// load html template
	Template, err = LoadTemplate()
	if err != nil {
		panic("failed to load html template")
	}

	// strip prefix from static fs
	sub, err := fs.Sub(StaticFiles, "static")
	if err != nil {
		panic("failed to strip 'static' from static files")
	}
	StaticFS = http.FS(sub)
}
