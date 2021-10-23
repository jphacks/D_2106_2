package database

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jphacks/D_2106_2/config"
)

type S3Client struct {
}

func NewS3Uploader() error {
	config, err := config.GetAwsConfig()
	if err != nil {
		return err
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
	fmt.Println(uploader)

	// f, err := os.Open(fileName + ".json")
	// if err != nil {
	// 	fmt.Errorf("failed to open file %q, %v", "test.json", err)
	// }

	// // Upload the file to S3.
	// res, err := uploader.Upload(&s3manager.UploadInput{
	// 	Bucket: aws.String("retweet-users"),
	// 	Key:    aws.String(fileName + ".json"),
	// 	Body:   f,
	// })

	// if err != nil {
	// 	fmt.Println(res)
	// 	if err, ok := err.(awserr.Error); ok && err.Code() == request.CanceledErrorCode {
	// 		fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
	// 	} else {
	// 		fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
	// 	}
	// 	os.Exit(1)
	// }

	return nil
}
