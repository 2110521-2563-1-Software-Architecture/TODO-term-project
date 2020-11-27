package main
import (
	"net/http"
	"service"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "online")
	})
	e.POST("/file",service.UploadFile())
	e.GET("/file",service.DownloadFile())
	e.Logger.Fatal(e.Start(":1323"))
}