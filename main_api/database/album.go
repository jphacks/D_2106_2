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

func (repo *AlbumRepository) StoreAlbum(album *domain.AlbumDB) (int, error) {
	result := repo.SqlHandler.Conn.Create(&album)
	if err := result.Error; err != nil {
		return -1, err
	}

	return album.Id, nil
}

func (repo *AlbumRepository) GetAllAlbums() ([]*domain.Album, error) {
	columns := []string{
		"albums.id",
		"user_id",
		"title",
		"started_at",
		"ended_at",
		"spot",
		"is_public",
		"images.url as ThumbnailImageUrl",
	}
	albums := []*domain.Album{}

	result := repo.SqlHandler.Conn.Table("albums").Select(columns).Joins("left join images on images.id = albums.thumbnail_image_id").Scan(&albums)
	if err := result.Error; err != nil {
		return nil, err
	}

	return albums, nil
}

func (repo *AlbumRepository) GetAlbumsByUsers(userId string) ([]*domain.Album, error) {
	columns := []string{
		"albums.id",
		"user_id",
		"title",
		"started_at",
		"ended_at",
		"spot",
		"is_public",
		"images.url as ThumbnailImageUrl",
	}
	join := "left join images on images.id = albums.thumbnail_image_id"
	albums := []*domain.Album{}

	result := repo.SqlHandler.Conn.Table("albums").Select(columns).Joins(join).Where("user_id = ?", userId).Scan(&albums)
	if err := result.Error; err != nil {
		return nil, err
	}

	return albums, nil
}
