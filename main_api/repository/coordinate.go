package repository

import "github.com/jphacks/D_2106_2/domain"

type CoordinateRepository interface {
	StoreCoordinates(coordinates []*domain.Coordinate) ([]int, error)
	GetCoordinateByImageId(imageId int) (*domain.Coordinate, error)
	GetCoordinatesByAlbumId(albumId int) ([]*domain.Coordinate, error)
}
