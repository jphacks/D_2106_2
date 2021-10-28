package testutils

import "github.com/jphacks/D_2106_2/domain"

type FakeAlbumRepository struct {
	FakeStoreAlbum       func(album *domain.Album) (int, error)
	FakeGetAllAlbums     func() ([]*domain.Album, error)
	FakeGetAlbumsByUsers func(user_id string) ([]*domain.Album, error)
}

func (repo FakeAlbumRepository) StoreAlbum(album *domain.Album) (int, error) {
	return repo.FakeStoreAlbum(album)
}

func (repo FakeAlbumRepository) GetAllAlbums() ([]*domain.Album, error) {
	return repo.FakeGetAllAlbums()
}

func (repo FakeAlbumRepository) GetAlbumsByUsers(user_id string) ([]*domain.Album, error) {
	return repo.FakeGetAlbumsByUsers(user_id)
}
