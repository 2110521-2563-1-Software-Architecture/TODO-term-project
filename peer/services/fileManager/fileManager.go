package fileManager

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const downloadFolder string = "files"
const path string = downloadFolder + "/"

func UploadFileToPeer(c *gin.Context) {

	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	if _, err := os.Stat(downloadFolder); os.IsNotExist(err) {
		os.Mkdir(downloadFolder, os.ModeDir)
	}

	if err := c.SaveUploadedFile(file, path+file.Filename); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})

}

func DownloadFileFromPeer(c *gin.Context) {

	name := c.Param("name")

	file, err := os.Open(path + name)
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

func Test(c *gin.Context) {
	name := c.Param("name")
	files := getFileFromPeer("test2.jpeg", "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.w3bai.com%2Fen-US%2Fcss%2Fcss3_images.html&psig=AOvVaw3GlAV4VZJN88FBbQHtUpKH&ust=1606821539407000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCJiyhOySqu0CFQAAAAAdAAAAABAD")
	// fmt.Printf("%v", files)
	c.JSON(http.StatusOK, gin.H{
		"message": files,
		"test":    name,
	})
}

// func fileIsValid(targetDir string, pattern []string) {

// 	for _, v := range pattern {
// 			matches, err := filepath.Glob(targetDir + v)

// 			if err != nil {
// 					fmt.Println(err)
// 			}

// 			if len(matches) != 0 {
// 					fmt.Println("Found : ", matches)
// 			}
// 	}
// }

func fileIsValid(filename string) bool {

	isValid := false
	filepath.Walk(path, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			if filename == f.Name() {
				isValid = true
				return nil
			}
		}
		return nil
	})
	return isValid
}

func getFileFromPeer(filename string, url string) error {

	//todo get tracker url

	filepath := path + filename

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// func updateTracker() {
// 	// Todo implement this function
// }
