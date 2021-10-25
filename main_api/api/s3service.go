package api

import (
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jphacks/D_2106_2/repository"
)

type S3service struct {
	S3Client
}

type UploadResult struct {
	Output *s3manager.UploadOutput
	Err    error
}

func NewS3service(client S3Client) repository.S3service {
	return &S3service{client}
}

func (client *S3service) S3Uploader(images []multipart.File, names []string) ([]string, error) {
	if client.Uploader == nil {
		return names, nil
	}

	c := make(chan UploadResult)
	imageUrls := make([]string, len(images))
	for i, image := range images {
		go func(i int, image multipart.File, c chan UploadResult) {
			output, err := client.Uploader.Upload(&s3manager.UploadInput{
				ACL:         aws.String("public-read"),
				Bucket:      aws.String(client.BucketName),
				Key:         aws.String("images/" + names[i]),
				Body:        image,
				ContentType: aws.String("image/jpeg"),
			})
			c <- UploadResult{Output: output, Err: err}
		}(i, image, c)
	}

	for range images {
		result := <-c
		if err := result.Err; err != nil {
			log.Print(err)
			return []string{}, nil
		}
		imageUrls = append(imageUrls, result.Output.Location)
	}

	return imageUrls, nil
}
