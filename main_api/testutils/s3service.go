package testutils

import "mime/multipart"

type FakeS3Service struct {
	FakeS3Uploader func(image []multipart.File, names []string) ([]string, error)
}

func (repo FakeS3Service) S3Uploader(image []multipart.File, names []string) ([]string, error) {
	return repo.FakeS3Uploader(image, names)
}
