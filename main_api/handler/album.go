package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/D_2106_2/usecase"
)

type AlbumHandler struct {
	uc usecase.AlbumUsecase
}

type Coordinate struct {
	Latitude  int
	Longitude int
	Timestamp string
}

type PostAlbumRequest struct {
	Locations []Coordinate
	Title     string
	StartAt   string
	EndAt     string
	IsPublic  bool
}

func NewAlbumHandler() *AlbumHandler {
	return &AlbumHandler{}
}

func (handler *AlbumHandler) GetAllAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "data"})
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

	c.JSON(http.StatusOK, gin.H{"data": "data"})
}
