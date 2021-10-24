package api

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jphacks/D_2106_2/config"
)

type S3Client struct {
}

func S3Uploader(images []multipart.File, names []string) (string, error) {
	config, err := config.GetAwsConfig()
	if err != nil {
		return "", err
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

	name := names[0]
	image := images[0]

	// Upload the file to S3.
	res, err := uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Bucket:      aws.String(config.S3BUCKET_MAME),
		Key:         aws.String("images/" + name),
		Body:        image,
		ContentType: aws.String("image/png"),
	})

	if err != nil {
		return "", err
	}

	imageUrl := res.Location

	return imageUrl, nil
}
