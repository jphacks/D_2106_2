package usecase

import (
	"mime/multipart"

	"github.com/jphacks/D_2106_2/repository"
)

type ImageUsecase struct {
	ImageRepo repository.ImageRepository
	CoordinateRepo repository.CoordinateRepository
}

func (uc *ImageUsecase) UploadImages(albumId int, images []multipart.File, names []string) error {

	/*
		1. upload image to s3, return image URI
		2. store image information to database
		3. 画面と座標を紐付ける
	*/

	coordinates, err :=uc.CoordinateRepo.GetCoordinatesByAlbumId(albumId)
	for image, _ := range images{

	}

	return nil
}
