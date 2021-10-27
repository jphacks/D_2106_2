package repository

import "github.com/jphacks/D_2106_2/domain"

type AlbumRepository interface {
	StoreAlbum(album *domain.Album) (int, error)
	GetAllAlbums() ([]*domain.Album, error)
	GetAlbumsByUsers(user_id string) ([]*domain.Album, error)
}
