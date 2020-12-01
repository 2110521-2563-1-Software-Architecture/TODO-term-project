package controllers

import (
	"net/http"
	s "tracker/services"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	service *s.FileService
}

type AddFileToPeerBody struct {
	PeerAddr string
	FileName string
}

func NewFileController(svc *s.FileService) *FileController {
	return &FileController{
		service: svc,
	}
}

func (con *FileController) GetAllFileNames(c *gin.Context) {
	result, err := con.service.GetAllFileNames()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "File names not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func (con *FileController) GetPeerWithFile(c *gin.Context) {
	fileName := c.Query("name")
	if fileName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file name",
		})
		return
	}
	result, err := con.service.GetPeersWithFile(fileName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "File name not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func (con *FileController) AddFileToPeer(c *gin.Context) {
	body := AddFileToPeerBody{}
	err := c.BindJSON(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body",
		})
		return
	}
	err = con.service.AddFileToPeer(body.FileName, body.PeerAddr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Track record added",
	})
}
