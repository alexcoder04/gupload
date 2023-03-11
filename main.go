package main

import (
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

	r.Run(":" + Port)
}
