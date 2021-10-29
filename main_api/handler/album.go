package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jphacks/D_2106_2/api"
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
	"github.com/jphacks/D_2106_2/usecase"
)

type AlbumHandler struct {
	uc usecase.AlbumUsecase
}

type GetAllAlbumsResponse struct {
	Albums []*domain.AlbumResponse `json:"albums"`
}

type PostAlbumRequest struct {
	Locations []*domain.Location `json:"locations"`
	UserId    string             `json:"userId"`
	Title     string             `json:"title"`
	StartAt   int64              `json:"startedAt"`
	EndAt     int64              `json:"endedAt"`
	IsPublic  bool               `json:"isPublic"`
}

type PostAlbumThumbnailRequest struct {
	AlbumId            int    `json:"albumId"`
	ThumbnailImageName string `json:"thumbnailImageName"`
}

type PostAlbumResponse struct {
	Id int `json:"id"`
}

func NewAlbumHandler(
	albumRepo repository.AlbumRepository,
	coordinateRepo repository.CoordinateRepository,
	imageRepo repository.ImageRepository,
) *AlbumHandler {
	uc := usecase.AlbumUsecase{AlbumRepo: albumRepo, CoordinateRepo: coordinateRepo, ImageRepo: imageRepo}

	return &AlbumHandler{uc: uc}
}

func (handler *AlbumHandler) GetAllAlbums(c *gin.Context) {
	albums, err := handler.uc.GetAllAlbums()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": FailedGetAlbum.Error()})
		return
	}

	response := make([]*domain.AlbumResponse, len(albums))
	for i, album := range albums {
		albumResponse := album.ToResponse()
		response[i] = albumResponse
	}

	c.JSON(http.StatusOK, gin.H{"data": GetAllAlbumsResponse{response}})
}

func (handler *AlbumHandler) GetUserAlbums(c *gin.Context) {
	userId := c.Query("user_id")
	if userId == "" {
		err := InvalidRequest
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums, err := handler.uc.GetUserAlbums(userId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": FailedGetAlbum.Error()})
		return
	}

	response := make([]*domain.AlbumResponse, len(albums))
	for i, album := range albums {
		albumResponse := album.ToResponse()
		response[i] = albumResponse
	}

	c.JSON(http.StatusOK, gin.H{"data": GetAllAlbumsResponse{response}})
}

func (handler *AlbumHandler) GetAlbumDetail(c *gin.Context) {
	albumId, _ := strconv.Atoi(c.Query("album_id"))
	lat1, _ := strconv.ParseFloat(c.Query("lat1"), 64)
	lon1, _ := strconv.ParseFloat(c.Query("lon1"), 64)
	lat2, _ := strconv.ParseFloat(c.Query("lat2"), 64)
	lon2, _ := strconv.ParseFloat(c.Query("lon2"), 64)

	if albumId <= 0 {
		err := InvalidRequest
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	latCondition := (-90 > lat1) || (lat2 > 90) || (lat1 >= lat2)
	lonCondition := (-180 > lon1) || (lon2 > 180) || (lon1 >= lon2)
	if latCondition || lonCondition {
		err := InvalidCoordinate
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clusterData, err := handler.uc.ClusteringGpsPoint(albumId, lat1, lat2, lon1, lon2)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if clusterData == nil {
		clusterData = &api.ClusterData{}
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
	responseData, err := handler.uc.ClusteringData2Response(albumId, &tempCoordinates)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": FailedClustering.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": *responseData})
}

func (handler *AlbumHandler) PostAlbum(c *gin.Context) {
	req := PostAlbumRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": InvalidRequest.Error()})
		return
	}

	if req.UserId == "" || len(req.Locations) <= 0 {
		err = InvalidRequest
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albumId, err := handler.uc.CreateNewAlbum(
		req.Locations,
		req.UserId,
		req.Title,
		req.StartAt,
		req.EndAt,
		req.IsPublic,
	)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": FailedCreateNewAlbum.Error()})
		return
	}

	res := &PostAlbumResponse{Id: albumId}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (handler *AlbumHandler) PostAlbumThumbnail(c *gin.Context) {

	req := PostAlbumThumbnailRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": InvalidRequest.Error()})
		return
	}

	if err := handler.uc.UpdateThumbnailAndSpot(req.AlbumId, req.ThumbnailImageName); err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": FailedUpdateThumbnailAndSpot.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"request": "success"})
}
