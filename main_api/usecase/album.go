package usecase

import (
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
)

type AlbumUsecase struct {
	AlbumRepo      repository.AlbumRepository
	CoordinateRepo repository.CoordinateRepository
	ImageRepo      repository.ImageRepository
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

	coordinates := make([]*domain.Coordinate, len(locations))
	for i, locate := range locations {
		coordinates[i] = &domain.Coordinate{
			AlbumId:   albumId,
			Timestamp: locate.Timestamp,
			Latitude:  locate.Latitude,
			Longitude: locate.Longitude,
		}
	}

	_, err = uc.CoordinateRepo.StoreCoordinates(coordinates)

	if err != nil {
		return -1, err
	}

	return albumId, nil
}

func (uc *AlbumUsecase) GetAllAlbums() ([]*domain.Album, error) {
	album, err := uc.AlbumRepo.GetAllAlbums()
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (uc *AlbumUsecase) GetUserAlbums(userId int) ([]*domain.Album, error) {
	album, err := uc.AlbumRepo.GetAlbumsByUsers(userId)
	if err != nil {
		return nil, err
	}

	return album, nil
}

// ClusteringGpsPoint clusters gps points
func (uc *AlbumUsecase) ClusteringGpsPoint(
	albumId int,
	latitudeMin float64,
	latitudeMax float64,
	longitudeMin float64,
	longitudeMax float64,
) (int, error) {
	var used_coordinates []domain.Coordinate
	images, err := uc.ImageRepo.GetImagesByAlbumId(albumId)
	// a, err := uc.AlbumRepo.GetCoordinatesByImageId
	for _, image := range images {
		coordinate, err := uc.CoordinateRepo.GetCoordinateByImageId(image.Id)
		if (latitudeMin < coordinate.Latitude) && (coordinate.Latitude < longitudeMax) {
			if (longitudeMin < coordinate.Longitude) && (coordinate.Longitude < longitudeMax) {
				used_coordinates = append(used_coordinates, *coordinate)
			}
		}
		if err != nil {
			return -1, err
		}
	}

	if err != nil {
		return -1, err
	}
	return -1, nil
}
