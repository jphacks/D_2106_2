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

func S3Uploader(images []multipart.File, names []string) ([]string, error) {
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

	// Upload the file to S3.
	c := make(chan *s3manager.UploadOutput)
	imageUrls := make([]string, len(images))
	for i, image := range images {
		go func(i int, image multipart.File, c chan *s3manager.UploadOutput) {
			res, _ := uploader.Upload(&s3manager.UploadInput{
				ACL:         aws.String("public-read"),
				Bucket:      aws.String(config.S3BUCKET_MAME),
				Key:         aws.String("images/" + names[i]),
				Body:        image,
				ContentType: aws.String("image/png"),
			})
			c <- res
		}(i, image, c)
	}

	for range images {
		result := <-c
		imageUrls = append(imageUrls, result.Location)
	}

	return imageUrls, nil
}
