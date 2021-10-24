package api

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jphacks/D_2106_2/repository"
)

type S3service struct {
	S3Client
}

func NewS3service(client S3Client) repository.S3service {
	return &S3service{client}
}

func (client *S3service) S3Uploader(images []multipart.File, names []string) ([]string, error) {
	c := make(chan *s3manager.UploadOutput)
	imageUrls := make([]string, len(images))
	for i, image := range images {
		go func(i int, image multipart.File, c chan *s3manager.UploadOutput) {
			res, _ := client.Uploader.Upload(&s3manager.UploadInput{
				ACL:         aws.String("public-read"),
				Bucket:      aws.String(client.BucketName),
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
