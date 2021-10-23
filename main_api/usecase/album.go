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

	coordinates := make([]*domain.Coordinate, len(locations), len(locations))
	for i, locate := range locations {
		coordinates[i] = &domain.Coordinate{
			AlbumId:   albumId,
			Timestamp: locate.Timestamp,
			Latitude:  locate.Latitude,
			Longitude: locate.Longitude,
		}
	}

	_, err = uc.CoordinateRepo.StoreCoordinates(coordinates)

	return albumId, nil
}

// ClusteringGpsPoint clusters gps points
func (uc *AlbumUsecase) ClusteringGpsPoint(
	albumId int,
	latitudeMin float64,
	latitudeMax float64,
	longitudeMin float64,
	longitudeMax float64,
) (int, error) {

	return 0, nil
}
