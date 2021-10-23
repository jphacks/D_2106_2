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
	var images []multipart.File
	var names []string
	var err error

	for i := 0; i < 5; i++ {
		filename := "image" + strconv.Itoa(i+1)
		image, header, err := c.Request.FormFile(filename)
		if err != nil {
			log.Print(err)
			c.JSON(500, gin.H{"err": err.Error()})
			return
		}
		images = append(images, image)
		names = append(names, header.Filename)
	}

	albumIdStr := c.PostForm("album_id")
	albumId, _ := strconv.Atoi(albumIdStr)

	err = handler.uc.UploadImages(albumId, images, names)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": "data"})
}
