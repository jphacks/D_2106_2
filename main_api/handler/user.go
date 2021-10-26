package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/jphacks/D_2106_2/repository"
	"github.com/jphacks/D_2106_2/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc usecase.UserUsecase
}

type RegisterUserRequest struct {
	Username     string `json:"username"`
	DeviceId     string `json:"deviceId"`
	Introduction string `json:"introduction"`
}

type RegisterUserRespose struct {
	UserId int `json:"user_id"`
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	uc := usecase.UserUsecase{UserRepo: userRepo}

	return &UserHandler{uc: uc}
}

func (handler *UserHandler) RegisterUser(c *gin.Context) {
	req := RegisterUserRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if req.Username == "" || req.DeviceId == "" {
		err = errors.New("username or deviceId field is null")
		log.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	userId, err := handler.uc.RegisterNewUser(req.Username, req.DeviceId, req.Introduction)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if userId == -1 {
		err = errors.New("register usesr failed")
		log.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := &RegisterUserRespose{UserId: userId}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (handler *UserHandler) GetUser(c *gin.Context) {
	deviceId := c.Query("device_id")
	user, err := handler.uc.GetUserByDeviceId(deviceId)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
