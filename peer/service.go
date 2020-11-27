package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/labstack/echo"
)

func UploadFileToPeer(c echo.Context) error{
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create files folder
	if _, err := os.Stat("files"); os.IsNotExist(err) {
		os.Mkdir("files", os.ModeDir)
	}

	// Destination
	dst, err := os.Create("files/"+file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("File %s uploaded successfully!", file.Filename))
}


func DownloadFileFromPeer(c echo.Context)error{
	name := c.QueryParam("name")
	return c.File("files/"+name)
}
