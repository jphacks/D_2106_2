package handler

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

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
		errorHandler(c, http.StatusBadRequest, "`album_id` field not found")
		return
	}
	albumId, err := strconv.Atoi(albumIdStr)
	if err != nil {
		errorHandler(c, http.StatusBadRequest, "`album_id` is invalid value")
		return
	}

	imageNumStr := c.PostForm("image_num")
	if !ok || imageNumStr == "" {
		errorHandler(c, http.StatusBadRequest, "`image_num` field not found")
		return
	}
	imageNum, err := strconv.Atoi(imageNumStr)
	if err != nil {
		errorHandler(c, http.StatusBadRequest, "`image_num` is invalid value")
		return
	}

	for i := 0; i < imageNum; i++ {
		filename := "image" + strconv.Itoa(i+1)
		image, header, err := c.Request.FormFile(filename)
		if err != nil {
			log.Print(err)
			c.JSON(500, gin.H{"err": err.Error()})
			return
		}

		if !validateImageName(strings.Split(header.Filename, ".")[0]) {
			errorHandler(c, http.StatusBadRequest, "invalid file name")
			return
		}

		images = append(images, image)
		names = append(names, fmt.Sprintf("%s-%s", albumIdStr, header.Filename))

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

func validateImageName(nameString string) bool {
	_, err := strconv.Atoi(nameString)
	if err != nil {
		log.Print(err)
		return false
	}
	return true
}

func errorHandler(c *gin.Context, code int, message string) {
	log.Print(message)
	c.JSON(code, gin.H{"err": message})
}
