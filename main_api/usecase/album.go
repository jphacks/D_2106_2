package usecase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jphacks/D_2106_2/api"
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
	"github.com/jphacks/D_2106_2/utils"
)

type AlbumUsecase struct {
	AlbumRepo      repository.AlbumRepository
	CoordinateRepo repository.CoordinateRepository
	ImageRepo      repository.ImageRepository
}

type ResponseLocation struct {
	Id        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	ImageUrls []string  `json:"imageUrls"`
}

type ResponseLocationData struct {
	Location []ResponseLocation `json:"location"`
	Route    []*domain.Location `json:"route"`
}

func (uc *AlbumUsecase) CreateNewAlbum(
	locations []*domain.Location,
	userId string,
	title string,
	startAt int64,
	endedAt int64,
	isPublic bool,
) (int, error) {
	album := &domain.AlbumDB{
		UserId:    userId,
		Title:     title,
		StartedAt: utils.UnixToTime(startAt),
		EndedAt:   utils.UnixToTime(endedAt),
		IsPublic:  isPublic,
	}

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
			Timestamp: utils.UnixToTime(locate.Timestamp),
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

func (uc *AlbumUsecase) GetUserAlbums(userId string) ([]*domain.Album, error) {
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

func (uc *AlbumUsecase) ClusteringData2Response(albumId int, tempCoordinates *[]domain.Coordinate) (*ResponseLocationData, error) {
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

	route, err := uc.CoordinateRepo.GetRouteByAlbumId(albumId)
	if err != nil {
		return nil, err
	}

	locationData := ResponseLocationData{
		Location: locationList,
		Route:    route,
	}

	return &locationData, nil
}

func (uc *AlbumUsecase) UpdateThumbnailAndSpot(albumId int, thumbnailImageName string) error {
	strAlbumId := strconv.Itoa(albumId)
	imgName := fmt.Sprintf("%s-%s", strAlbumId, thumbnailImageName)
	image, err := uc.ImageRepo.GetImageByImageName(imgName)
	if err != nil {
		return err
	}
	coordinate, err := uc.CoordinateRepo.GetCoordinateById(image.CoordinateId)
	if err != nil {
		return err
	}
	city, prefecture, err := utils.GetMunicipalitiesByGeoLocation(coordinate.Latitude, coordinate.Longitude)
	if err != nil {
		return err
	}
	spot := fmt.Sprintf("%s%s", prefecture, city)

	if err := uc.AlbumRepo.UpdateThumbnailAndSpotByAlbumId(albumId, image.Id, spot); err != nil {
		return err
	}
	return nil
}
