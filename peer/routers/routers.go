package routers

import (
	"net/http"

	"peer/services/fileManager"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "online",
		})
	})

	r.POST("/uploadFile", fileManager.UploadFileToPeer)
	r.GET("/downloadFile/:name", fileManager.DownloadFileFromPeer)
	return r
}
