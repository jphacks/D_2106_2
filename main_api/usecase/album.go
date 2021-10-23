package usecase

import (
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
)

type AlbumUsecase struct {
	AlbumRepo      repository.AlbumRepository
	CoordinateRepo repository.CoordinateRepository
}

func (uc *AlbumUsecase) CreateNewAlbum(
	locations []*domain.Location,
	userId int,
	title string,
	startAt string,
	endedAt string,
	isPublic bool,
) (int, error) {
	album := &domain.Album{
		UserId:    userId,
		Title:     title,
		StartedAt: startAt,
		EndedAt:   endedAt,
		IsPublic:  isPublic,
	}

	/* TODO: implement create album */

	albumId, err := uc.AlbumRepo.StoreAlbum(album)
	if err != nil {
		return -1, err
	}

	// coordinateIds, err := uc.CoordinateRepo.StoreCoordinates(locations)

	return albumId, nil
}
