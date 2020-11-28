package main
import (
	"net/http"
	"github.com/labstack/echo"
)

func fileHandler() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "online")
	})
	e.POST("/file",UploadFileToPeer)
	e.GET("/file",DownloadFileFromPeer)
	e.Logger.Fatal(e.Start(":1323"))
}