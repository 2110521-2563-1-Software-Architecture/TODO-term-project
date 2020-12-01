package routers

import (
	"net/http"
	"tracker/controllers"
	"tracker/services"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

	// Initialize file routers
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	fileService := services.NewFileService(rdb)
	fileController := controllers.NewFileController(fileService)

	r.GET("/file", fileController.GetPeerWithFile)
	r.POST("/file", fileController.AddFileToPeer)
	return r
}
