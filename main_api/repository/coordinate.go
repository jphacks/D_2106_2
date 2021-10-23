package repository

import "github.com/jphacks/D_2106_2/domain"

type CoordinateRepository interface {
	StoreCoordinates(coordinates []*domain.Coordinate) ([]int, error)
}
