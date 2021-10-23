package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
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
	// albumId := c.Query("album_id")
	// lat1 := c.Query("lat1")
	// lon1 := c.Query("lon1")
	// lat2 := c.Query("lat2")
	// lon2 := c.Query("lon2")

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
