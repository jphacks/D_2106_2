package repository

import "github.com/jphacks/D_2106_2/domain"

type ImageRepository interface {
	GetImagesByAlbumId(albumId int) ([]*domain.Image, error)
	GetImagesByCoordinateId(coordinateId int) ([]*domain.Image, error)
	StoreImages(images []*domain.Image) ([]int, error)
	GetImagesById(imageId int) (*domain.Image, error)
}
