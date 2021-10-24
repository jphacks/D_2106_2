package database

import (
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
)

type ImageRepository struct {
	SqlHandler
}

func NewImageRepository(sqlHandler SqlHandler) repository.ImageRepository {
	return &ImageRepository{sqlHandler}
}

func (repo *ImageRepository) GetImagesByAlbumId(albumId int) ([]*domain.Image, error) {
	images := []*domain.Image{}
	result := repo.SqlHandler.Conn.Where("album_id = ?", albumId).Find(&images)
	if err := result.Error; err != nil {
		return nil, err
	}

	return images, nil
}

func (repo *ImageRepository) StoreImages(images []*domain.Image) ([]int, error) {
	result := repo.SqlHandler.Conn.Create(&images)
	if err := result.Error; err != nil {
		return nil, err
	}

	var idList []int
	for _, image := range images {
		idList = append(idList, image.Id)
	}

	return idList, nil
}

func (repo *ImageRepository) GetImagesByCoordinateId(coordinateId int) ([]*domain.Image, error) {
	images := []*domain.Image{}
	result := repo.SqlHandler.Conn.Where("coordinate_id = ?", coordinateId).Find(&images)
	if err := result.Error; err != nil {
		return nil, err
	}
  
	return images, nil
}
