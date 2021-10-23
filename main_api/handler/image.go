package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
}

func NewImageHandler() *ImageHandler {
	return &ImageHandler{}
}

func (handler *ImageHandler) UploadImages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "data"})
}
