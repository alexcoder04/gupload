package main

import (
	"html/template"
	"os"
)

type GuploadConfig struct {
	IP       string
	Hostname string
	Port     string

	Directory string

	Template *template.Template
}

var Config = LoadConfig()

func LoadConfig() GuploadConfig {
	ip, err := GetLocalIP()
	if err != nil {
		panic("failed to get local ip address")
	}

	// TODO port from environ

	host, err := os.Hostname()
	if err != nil {
		host = "unknown-pc"
	}

	// TODO shared dir

	tmpl, err := LoadTemplate()
	if err != nil {
		panic("failed to load html template")
	}

	return GuploadConfig{
		IP:       ip,
		Hostname: host,
		Port:     "1234",
		Template: tmpl,
	}
}
