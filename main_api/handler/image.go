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

func NewImageHandler(imageRepo repository.ImageRepository, s3service repository.S3service, coordinateRepo repository.CoordinateRepository) *ImageHandler {
	uc := usecase.ImageUsecase{ImageRepo: imageRepo, S3service: s3service, CoordinateRepo: coordinateRepo}

	return &ImageHandler{uc: uc}
}

func (handler *ImageHandler) UploadImages(c *gin.Context) {
	var images []multipart.File
	var names []string
	var err error

	albumIdStr, ok := c.GetPostForm("album_id")
	if !ok || albumIdStr == "" {
		message := "`album_id` field not found"
		log.Print(message)
		c.JSON(400, gin.H{"err": message})
		return
	}
	albumId, _ := strconv.Atoi(albumIdStr)

	imageNumStr := c.PostForm("image_num")
	if !ok || imageNumStr == "" {
		message := "`image_num` field not found"
		log.Print(message)
		c.JSON(400, gin.H{"err": message})
		return
	}
	imageNum, _ := strconv.Atoi(imageNumStr)

	for i := 0; i < imageNum; i++ {
		filename := "image" + strconv.Itoa(i+1)
		image, header, err := c.Request.FormFile(filename)
		if err != nil {
			log.Print(err)
			c.JSON(500, gin.H{"err": err.Error()})
			return
		}
		images = append(images, image)
		names = append(names, header.Filename)

		log.Printf("Uploade %s, Size: %d", header.Filename, header.Size)
	}

	err = handler.uc.UploadImages(albumId, images, names)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "data"})
}
