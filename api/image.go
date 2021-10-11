package api

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"

	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/infrastructure"
	"github.com/gin-gonic/gin"
)

type uploadImageRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// Upload upload files.
func (server *Server) uploadImage(ctx *gin.Context) {
	var req uploadImageRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// Create S3 session
	awsS3 := infrastructure.NewAwsS3()

	form, _ := ctx.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		fileName := file.Filename
		if fileName == "" {
			err := errors.New("fileName is required")
			ctx.JSON(400, gin.H{"message": err.Error()})
		}
		re1 := regexp.MustCompile("[^`><{}][()#%~|&$@=;: +,?\\\\]")
		fileName = re1.ReplaceAllString(fileName, "")
		ext := filepath.Ext(fileName)
		uploadFile, err := file.Open()
		if err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
		}
		// Upload the file to S3.
		url, err := awsS3.Upload(uploadFile, fileName, ext)
		if err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
		}
		fmt.Printf("%+v url", url)
		arg := db.UpdateUserImageParams{
			ID:    req.ID,
			Image: url,
		}
		user, err := server.store.UpdateUserImage(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
