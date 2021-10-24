package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
	"github.com/jphacks/D_2106_2/usecase"
)

type AlbumHandler struct {
	uc usecase.AlbumUsecase
}

type PostAlbumRequest struct {
	Locations []*domain.Location `json:"locations"`
	UserId    int                `json:"userId"`
	Title     string             `json:"title"`
	StartAt   string             `json:"startedAt"`
	EndAt     string             `json:"endedAt"`
	IsPublic  bool               `json:"isPublic"`
}

type PostAlbumResponse struct {
	Id int `json:"id"`
}

// type TempLocation struct {
// 	GpsId     int     `json:"id"`
// 	Latitude  float64 `json:"latitude`
// 	Longitude float64 `json:"longitude`
// }

func NewAlbumHandler(albumRepo repository.AlbumRepository, coordinateRepo repository.CoordinateRepository, imageRepo repository.ImageRepository) *AlbumHandler {
	uc := usecase.AlbumUsecase{AlbumRepo: albumRepo, CoordinateRepo: coordinateRepo, ImageRepo: imageRepo}

	return &AlbumHandler{uc: uc}
}

func (handler *AlbumHandler) GetAllAlbums(c *gin.Context) {
	albums, err := handler.uc.GetAllAlbums()
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": albums})
}

func (handler *AlbumHandler) GetUserAlbums(c *gin.Context) {
	userIdStr := c.Query("album_id")
	userId, _ := strconv.Atoi(userIdStr)
	albums, err := handler.uc.GetUserAlbums(userId)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": albums})
}

func (handler *AlbumHandler) GetAlbum(c *gin.Context) {
	// userId := c.Query("user_id")

	c.JSON(http.StatusOK, gin.H{"data": "data"})
}

func (handler *AlbumHandler) GetAlbumDetail(c *gin.Context) {
	albumId, _ := strconv.Atoi(c.Query("album_id"))
	lat1, _ := strconv.ParseFloat(c.Query("lat1"), 64)
	lon1, _ := strconv.ParseFloat(c.Query("lon1"), 64)
	lat2, _ := strconv.ParseFloat(c.Query("lat2"), 64)
	lon2, _ := strconv.ParseFloat(c.Query("lon2"), 64)
	if albumId <= 0 {
		c.JSON(400, gin.H{"err": "album_id is invalid"})
		return
	}
	if (-90 > lat1) || (lat2 > 90) || (lat1 >= lat2) {
		c.JSON(400, gin.H{"err": "latitude is really? -90 = lat = 90"})
		return
	}
	if (-180 > lon1) || (lon2 > 180) || (lon1 >= lon2) {
		c.JSON(400, gin.H{"err": "longitude is really? -180 = lat = 180"})
		return
	}
	clusterData, err := handler.uc.ClusteringGpsPoint(albumId, lat1, lat2, lon1, lon2)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	var tempCoordinates []domain.Coordinate
	for _, data := range clusterData.ClusterData {
		meanLatitude := data.MeanLatitude
		meanLongitude := data.MeanLongitude
		for _, gpsId := range data.GpsIdBelongsTo {
			tempCoordinates = append(tempCoordinates, domain.Coordinate{
				Id:        gpsId,
				Latitude:  meanLatitude,
				Longitude: meanLongitude,
			})
		}
	}
	responseData, err := handler.uc.ClusteringData2Response(&tempCoordinates)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": *responseData})
}

func (handler *AlbumHandler) PostAlbum(c *gin.Context) {
	req := PostAlbumRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	albumId, err := handler.uc.CreateNewAlbum(req.Locations, req.UserId, req.Title, req.StartAt, req.EndAt, req.IsPublic)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	res := &PostAlbumResponse{Id: albumId}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
