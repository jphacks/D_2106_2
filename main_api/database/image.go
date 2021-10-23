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
