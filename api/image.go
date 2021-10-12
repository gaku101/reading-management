package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"

	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/infrastructure"
	"github.com/gin-gonic/gin"
)

type uploadImageRequest struct {
	Username string `uri:"username" binding:"required,min=1"`
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
		fileName = strings.ReplaceAll(fileName, " ", "")
		re1 := regexp.MustCompile("[^`><{}][()#%~|&$@=;: +,?\\\\]")
		fileName = re1.ReplaceAllString(fileName, "")
		fmt.Printf("fileNmae %+v", fileName)
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
		fmt.Printf("url %+v ", url)
		image, err := server.store.GetUserImage(ctx, req.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		if url != image {
			err = awsS3.Delete(image)
			if err != nil {
				ctx.JSON(400, gin.H{"message": err.Error()})
			}
		}
		arg := db.UpdateUserImageParams{
			Username: req.Username,
			Image:    url,
		}
		user, err := server.store.UpdateUserImage(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		rsp := newUserResponse(user)
		ctx.JSON(http.StatusOK, rsp)
	}
}
