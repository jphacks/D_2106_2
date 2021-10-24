package usecase

import (
	"fmt"

	"github.com/jphacks/D_2106_2/api"
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
)

type AlbumUsecase struct {
	AlbumRepo      repository.AlbumRepository
	CoordinateRepo repository.CoordinateRepository
	ImageRepo      repository.ImageRepository
}

type ResponseLocation struct {
	Id        int
	Timestamp string
	Latitude  float64
	Longitude float64
	ImageUrls []string
}

type ResponseLocationData struct {
	Location []ResponseLocation
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
		isShow := false
		if i%10 == 0 || i+1 == len(locations) {
			isShow = true
		}
		coordinates[i] = &domain.Coordinate{
			AlbumId:   albumId,
			Timestamp: locate.Timestamp,
			Latitude:  locate.Latitude,
			Longitude: locate.Longitude,
			IsShow:    isShow,
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
) (*api.ClusterData, error) {
	var used_coordinates []domain.Coordinate
	images, err := uc.ImageRepo.GetImagesByAlbumId(albumId)
	if err != nil {
		return nil, err
	}
	for _, image := range images {
		coordinate, err := uc.CoordinateRepo.GetCoordinateById(image.CoordinateId)
		if err != nil {
			return nil, err
		}
		if (latitudeMin < coordinate.Latitude) && (coordinate.Latitude < latitudeMax) {
			if (longitudeMin < coordinate.Longitude) && (coordinate.Longitude < longitudeMax) {
				fmt.Println(latitudeMin, coordinate.Latitude, latitudeMax)
				used_coordinates = append(used_coordinates, *coordinate)
			}
		}
	}
	if len(used_coordinates) == 0 {
		return nil, nil
	}
	gpsData := api.Coordinates2GpsData(used_coordinates)
	clusterData, err := api.GetClusteringApi(gpsData)
	if err != nil {
		return nil, err
	}
	return clusterData, nil
}

func (uc *AlbumUsecase) ClusteringData2Response(tempCoordinates *[]domain.Coordinate) (*ResponseLocationData, error) {
	var locationList []ResponseLocation
	for _, tempCoordinate := range *tempCoordinates {
		coordinate, err := uc.CoordinateRepo.GetCoordinateById(tempCoordinate.Id)
		if err != nil {
			return nil, err
		}
		images, err := uc.ImageRepo.GetImagesByCoordinateId(tempCoordinate.Id)
		if err != nil {
			return nil, err
		}
		var imageUrls []string
		for _, image := range images {
			imageUrls = append(imageUrls, image.Url)
		}
		locationList = append(locationList, ResponseLocation{
			Id:        tempCoordinate.Id,
			Timestamp: coordinate.Timestamp,
			Latitude:  tempCoordinate.Latitude,
			Longitude: tempCoordinate.Longitude,
			ImageUrls: imageUrls,
		})
	}
	locationData := ResponseLocationData{
		Location: locationList,
	}
	return &locationData, nil
}
