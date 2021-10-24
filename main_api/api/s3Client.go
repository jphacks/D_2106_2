package api

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jphacks/D_2106_2/config"
)

type S3Client struct {
	BucketName string
	Uploader   *s3manager.Uploader
}

func NewS3client() (*S3Client, error) {
	config, err := config.GetAwsConfig()
	if err != nil {
		return nil, err
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			config.ACCESS_KEY,
			config.SECRET_KEY,
			"",
		),
		Region: aws.String(config.REGION),
	}))

	uploader := s3manager.NewUploader(sess)

	return &S3Client{BucketName: config.S3BUCKET_MAME, Uploader: uploader}, nil
}
