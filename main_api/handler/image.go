package handler

import (
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/D_2106_2/repository"
	"github.com/jphacks/D_2106_2/usecase"
)

type ImageHandler struct {
	uc usecase.ImageUsecase
}

func NewImageHandler(imageRepo repository.ImageRepository) *ImageHandler {
	uc := usecase.ImageUsecase{ImageRepo: imageRepo}

	return &ImageHandler{uc: uc}
}

func (handler *ImageHandler) UploadImages(c *gin.Context) {
	image1, _, err := c.Request.FormFile("image1")
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
	}

	var images []multipart.File
	images = append(images, image1)

	albumIdStr := c.PostForm("album_id")
	albumId, _ := strconv.Atoi(albumIdStr)

	err = handler.uc.UploadImages(albumId, images)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": "data"})
}
