package fileManager

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadFileToPeer(c *gin.Context) {
	file, err := c.FormFile("file")

	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// src, err := file.Open()
	// if err != nil {
	// 	return err
	// }
	// defer src.Close()

	// Create files folder
	if _, err := os.Stat("files"); os.IsNotExist(err) {
		os.Mkdir("files", os.ModeDir)
	}

	// // Destination
	// dst, err := os.Create("files/" + file.Filename)
	// if err != nil {
	// 	return err
	// }
	// defer dst.Close()

	// // Copy
	// if _, err = io.Copy(dst, src); err != nil {
	// 	return err
	// }

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, "files/"+file.Filename); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})

}

func DownloadFileFromPeer(c *gin.Context) {
	name := c.Param("name")
	file, err := os.Open("files/" + name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "File not found",
		})
		return
	}
	defer file.Close()

	c.Writer.Header().Add("Content-type", "application/octet-stream")
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "File download fail.",
		})
		return
	}
}
