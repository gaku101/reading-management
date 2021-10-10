package api

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/gaku101/my-portfolio/infrastructure"
	"github.com/gin-gonic/gin"
)

// Upload upload files.
func (server *Server) uploadImage(c *gin.Context) {
	// Create S3 session
	awsS3 := infrastructure.NewAwsS3()

	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		fileName := file.Filename
		if fileName == "" {
			err := errors.New("fileName is required")
			c.JSON(400, gin.H{"message": err.Error()})
		}
		ext := filepath.Ext(fileName)
		fmt.Printf("ext %+v", ext)
		uploadFile, err := file.Open()
		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
		}
		// Upload the file to S3.
		url, err := awsS3.Upload(uploadFile, fileName, ext)
		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
		}
		fmt.Printf("%+v url", url)
		c.JSON(201, gin.H{"message": "success"})
	}
}
