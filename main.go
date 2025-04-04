package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// static files
	r.StaticFS("/static", StaticFS)
	r.GET("/qrcode.png", HandlerQRCode)

	// main page
	r.GET("/", HandlerIndex)
	r.POST("/", HandlerUpload)

	// download files
	r.GET("/download/:path", HandlerDownload)
	r.GET("/zip/:path", HandlerZip)

	fmt.Printf("Starting server on %s:%s.\n", Config.IP, Config.Port)
	r.Run(":" + Config.Port)
}
