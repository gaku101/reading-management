package infrastructure

import (
	"fmt"
	"log"

	"github.com/gaku101/my-portfolio/util"
)

type Config struct {
	Aws struct {
		S3 struct {
			Region string
			Bucket string
			AccessKeyID string
			SecretAccessKey string
		}
	}
}

func NewConfig() *Config {

	c := new(Config)
	appConfig, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	fmt.Println("AccessKeyID\n", appConfig.AwsAccessKeyID)
	fmt.Println("SecretAccessKey\n", appConfig.AwsSecretAccessKey)
	c.Aws.S3.Region = "ap-northeast-1"
	c.Aws.S3.Bucket = "my-portfolio-bucket-01"
	c.Aws.S3.AccessKeyID = appConfig.AwsAccessKeyID
	c.Aws.S3.SecretAccessKey = appConfig.AwsSecretAccessKey

	return c
}
