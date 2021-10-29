package testutils

import "github.com/jphacks/D_2106_2/domain"

type FakeImageRepository struct {
	FakeGetImagesByAlbumId      func(albumId int) ([]*domain.Image, error)
	FakeGetImagesByCoordinateId func(coordinateId int) ([]*domain.Image, error)
	FakeStoreImages             func(images []*domain.Image) ([]int, error)
	FakeGetImagesById           func(imageId int) (*domain.Image, error)
	FakeGetImageByImageName     func(imageName string) (*domain.Image, error)
}

func (repo FakeImageRepository) GetImagesByAlbumId(albumId int) ([]*domain.Image, error) {
	return repo.FakeGetImagesByAlbumId(albumId)
}

func (repo FakeImageRepository) GetImagesByCoordinateId(coordinateId int) ([]*domain.Image, error) {
	return repo.FakeGetImagesByCoordinateId(coordinateId)
}

func (repo FakeImageRepository) StoreImages(images []*domain.Image) ([]int, error) {
	return repo.FakeStoreImages(images)
}

func (repo FakeImageRepository) GetImagesById(imageId int) (*domain.Image, error) {
	return repo.FakeGetImagesById(imageId)
}

func (repo FakeImageRepository) GetImageByImageName(imageName string) (*domain.Image, error) {
	return repo.FakeGetImageByImageName(imageName)
}
