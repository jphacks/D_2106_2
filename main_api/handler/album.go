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

func NewAlbumHandler(albumRepo repository.AlbumRepository, coordinateRepo repository.CoordinateRepository) *AlbumHandler {
	uc := usecase.AlbumUsecase{AlbumRepo: albumRepo, CoordinateRepo: coordinateRepo}

	return &AlbumHandler{uc: uc}
}

func (handler *AlbumHandler) GetAllAlbums(c *gin.Context) {
	albums, err := handler.uc.GetAllAlbums()
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
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
	handler.uc.ClusteringGpsPoint(albumId, lat1, lat2, lon1, lon2)
	c.JSON(http.StatusOK, gin.H{"data": "data"})
}

func (handler *AlbumHandler) PostAlbum(c *gin.Context) {
	req := PostAlbumRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
	}

	albumId, err := handler.uc.CreateNewAlbum(req.Locations, req.UserId, req.Title, req.StartAt, req.EndAt, req.IsPublic)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
	}

	res := &PostAlbumResponse{Id: albumId}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
