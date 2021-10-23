package database

import (
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
)

type CoordinateRepository struct {
	SqlHandler
}

func NewCoordinateRepository(sqlHandler SqlHandler) repository.CoordinateRepository {
	return &CoordinateRepository{sqlHandler}
}

func (repo *CoordinateRepository) StoreCoordinates(coordinates []*domain.Coordinate) ([]int, error) {
	result := repo.SqlHandler.Conn.Create(&coordinates)
	if err := result.Error; err != nil {
		return nil, err
	}

	var idList []int
	for _, coordinate := range coordinates {
		idList = append(idList, coordinate.Id)
	}

	return idList, nil
}
