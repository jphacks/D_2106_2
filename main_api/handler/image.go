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

	albumIdStr := c.PostForm("album_id")
	albumId, _ := strconv.Atoi(albumIdStr)

	imageNumStr := c.PostForm("image_num")
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

		// x, err := exif.Decode(image)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(x)
	}

	err = handler.uc.UploadImages(albumId, images, names)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": "data"})
}
