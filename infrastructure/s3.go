package infrastructure

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AwsS3 struct {
	Config   *Config
	Keys     AwsS3URLs
	Uploader *s3manager.Uploader
	Svc *s3.S3
}

type AwsS3URLs struct {
	Folder string
}

func NewAwsS3() *AwsS3 {
	config := NewConfig()
	// s3manager.Uploader を初期化
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(config.Aws.S3.Region)},
		Profile: "default",
	})
	if err != nil {
		panic(err)
	}

	return &AwsS3{
		Config: config,
		Keys: AwsS3URLs{
			Folder: "images",
		},
		// Create an uploader with the session and default options
		Uploader: s3manager.NewUploader(sess),
		Svc: s3.New(sess),
	}
}

func (a *AwsS3) Upload(file multipart.File, fileName string, extension string) (url string, err error) {

	if fileName == "" {
		return "", errors.New("fileName is required")
	}

	var contentType string

	switch extension {
	case ".jpg":
		contentType = "image/jpeg"
	case ".jpeg":
		contentType = "image/jpeg"
	case ".gif":
		contentType = "image/gif"
	case ".png":
		contentType = "image/png"
	default:
		return "", errors.New("this extension is invalid")
	}

	// Upload the file to S3.
	result, err := a.Uploader.Upload(&s3manager.UploadInput{
		ACL:    aws.String("public-read"),
		Body:   file,
		Bucket: aws.String(a.Config.Aws.S3.Bucket),
		ContentType: aws.String(contentType),
		Key:         aws.String(a.Keys.Folder + "/" + fileName),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}
	return result.Location, nil
}
