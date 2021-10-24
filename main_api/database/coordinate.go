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

func (repo *CoordinateRepository) GetCoordinateByImageId(imageId int) (*domain.Coordinate, error) {
	var coordinate *domain.Coordinate
	result := repo.SqlHandler.Conn.Where("image_id = ?", imageId).First(&coordinate)
	if err := result.Error; err != nil {
		return nil, err
	}
	return coordinate, nil
}

func (repo *CoordinateRepository) GetCoordinateById(coordinateId int) (*domain.Coordinate, error) {
	var coordinate *domain.Coordinate
	result := repo.SqlHandler.Conn.Where("id = ?", coordinateId).First(&coordinate)
	if err := result.Error; err != nil {
		return nil, err
	}
	return coordinate, nil
}


func (repo *CoordinateRepository) GetCoordinatesByAlbumId(albumId int) ([]*domain.Coordinate, error) {
	coordinates := []*domain.Coordinate{}
	result := repo.SqlHandler.Conn.Where("album_id = ?", imagealbumIdId).Order("timestamp asc").Find(&coordinates)
	if err := result.Error; err != nil {
		return nil, err
	}

	return coordinates, nil
}
