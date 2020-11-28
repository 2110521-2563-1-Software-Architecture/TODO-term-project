package main

import (
	"github.com/labstack/echo"
)

func fileHandler() {
	e := echo.New()

	e.POST("/file", UploadFileToPeer)
	e.GET("/file", DownloadFileFromPeer)
	e.Logger.Fatal(e.Start(":9000"))
}
