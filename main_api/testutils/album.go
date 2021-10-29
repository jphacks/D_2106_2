package testutils

import "github.com/jphacks/D_2106_2/domain"

type FakeAlbumRepository struct {
	FakeStoreAlbum                      func(album *domain.AlbumDB) (int, error)
	FakeGetAllAlbums                    func() ([]*domain.Album, error)
	FakeGetAlbumsByUsers                func(user_id string) ([]*domain.Album, error)
	FakeUpdateThumbnailAndSpotByAlbumId func(albumId int, thumbnailImageId int, spot string) error
}

func (repo FakeAlbumRepository) StoreAlbum(album *domain.AlbumDB) (int, error) {
	return repo.FakeStoreAlbum(album)
}

func (repo FakeAlbumRepository) GetAllAlbums() ([]*domain.Album, error) {
	return repo.FakeGetAllAlbums()
}

func (repo FakeAlbumRepository) GetAlbumsByUsers(user_id string) ([]*domain.Album, error) {
	return repo.FakeGetAlbumsByUsers(user_id)
}

func (repo FakeAlbumRepository) UpdateThumbnailAndSpotByAlbumId(albumId int, thumbnailImageId int, spot string) error {
	return repo.FakeUpdateThumbnailAndSpotByAlbumId(albumId, thumbnailImageId, spot)
}
