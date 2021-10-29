package testutils

import "github.com/jphacks/D_2106_2/domain"

type FakeCoordinateRepository struct {
	FakeStoreCoordinates        func(coordinates []*domain.Coordinate) ([]int, error)
	FakeGetCoordinateByImageId  func(imageId int) (*domain.Coordinate, error)
	FakeGetCoordinatesByAlbumId func(albumId int) ([]*domain.Coordinate, error)
	FakeGetCoordinateById       func(coordinateId int) (*domain.Coordinate, error)
	FakeGetRouteByAlbumId       func(albumId int) ([]*domain.Location, error)
}

func (repo FakeCoordinateRepository) StoreCoordinates(coordinates []*domain.Coordinate) ([]int, error) {
	return repo.FakeStoreCoordinates(coordinates)
}

func (repo FakeCoordinateRepository) GetCoordinateByImageId(imageId int) (*domain.Coordinate, error) {
	return repo.FakeGetCoordinateByImageId(imageId)
}

func (repo FakeCoordinateRepository) GetCoordinatesByAlbumId(albumId int) ([]*domain.Coordinate, error) {
	return repo.FakeGetCoordinatesByAlbumId(albumId)
}

func (repo FakeCoordinateRepository) GetCoordinateById(coordinateId int) (*domain.Coordinate, error) {
	return repo.FakeGetCoordinateById(coordinateId)
}

func (repo FakeCoordinateRepository) GetRouteByAlbumId(albumId int) ([]*domain.Location, error) {
	return repo.FakeGetRouteByAlbumId(albumId)
}
