package repository

import "mime/multipart"

type S3service interface {
	S3Uploader(images []multipart.File, names []string) ([]string, error)
}
