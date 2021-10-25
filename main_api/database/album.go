package database

import (
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
)

type AlbumRepository struct {
	SqlHandler
}

func NewAlbumRepository(sqlHandler SqlHandler) repository.AlbumRepository {
	return &AlbumRepository{sqlHandler}
}

func (repo *AlbumRepository) StoreAlbum(album *domain.Album) (int, error) {
	result := repo.SqlHandler.Conn.Create(&album)
	if err := result.Error; err != nil {
		return -1, err
	}

	return album.Id, nil
}

func (repo *AlbumRepository) GetAllAlbums() ([]*domain.Album, error) {
	albums := []*domain.Album{}
	result := repo.SqlHandler.Conn.Find(&albums)
	if err := result.Error; err != nil {
		return nil, err
	}

	return albums, nil
}

func (repo *AlbumRepository) GetAlbumsByUsers(userId int) ([]*domain.Album, error) {
	albums := []*domain.Album{}
	result := repo.SqlHandler.Conn.Where("user_id = ?", userId).Find(&albums)
	if err := result.Error; err != nil {
		return nil, err
	}

	return albums, nil
}
