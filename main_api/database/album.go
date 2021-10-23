package database

import (
	"fmt"

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
	fmt.Println(album)
	result := repo.SqlHandler.Conn.Create(&album)
	if err := result.Error; err != nil {
		return -1, nil
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
