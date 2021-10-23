package usecase

import (
	"mime/multipart"

	"github.com/jphacks/D_2106_2/repository"
)

type ImageUsecase struct {
	ImageRepo repository.ImageRepository
}

func (uc *ImageUsecase) UploadImages(albumId int, images []multipart.File) error {
	/*
		1. upload image to s3, return image URI
		2. store image information to database
	*/

	return nil
}
