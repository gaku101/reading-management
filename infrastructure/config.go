package infrastructure

type Config struct {
	Aws struct {
		S3 struct {
			Region          string
			Bucket          string
		}
	}
}

func NewConfig() *Config {

	c := new(Config)

	c.Aws.S3.Region = "ap-northeast-1"
	c.Aws.S3.Bucket = "my-portfolio-bucket-01"

	return c
}
