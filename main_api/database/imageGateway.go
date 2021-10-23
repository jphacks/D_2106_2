package database

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Client struct {
}

func NewS3Uploader() error {
	// TODO: upload image to s3
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("ap-northeast-1")},
		Profile: "default",
	})
	if err != nil {
		return err
	}

	uploader := s3manager.NewUploader(sess)

	data := strings.NewReader(`{"message": "hello world"}`)
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String("bucket-name"),
		Body:        aws.ReadSeekCloser(data),
		Key:         aws.String("path/to/file"),
		ContentType: aws.String("application/json"),
	})
	fmt.Println(output)
	if err != nil {
		return err
	}

	return nil
}
