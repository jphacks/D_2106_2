package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/jphacks/D_2106_2/repository"
	"github.com/jphacks/D_2106_2/usecase"

	"github.com/gin-gonic/gin"
)

var (
	InvalidRequest     = errors.New("Invalid reqquest parameter")
	FieldIsNull        = errors.New("username or deviceId field is null")
	FailedRegisterUser = errors.New("register usesr failed")
)

type UserHandler struct {
	uc usecase.UserUsecase
}

type RegisterUserRequest struct {
	DeviceId     string `json:"deviceId"`
	Username     string `json:"username"`
	Introduction string `json:"introduction"`
}

type RegisterUserRespose struct {
	UserId string `json:"user_id"`
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
		c.JSON(http.StatusBadRequest, gin.H{"error": InvalidRequest.Error()})
		return
	}

	if req.Username == "" || req.DeviceId == "" {
		err = FieldIsNull
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := handler.uc.RegisterNewUser(req.Username, req.DeviceId, req.Introduction)
	if err != nil {
		err = FailedRegisterUser
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := &RegisterUserRespose{UserId: userId}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (handler *UserHandler) GetUser(c *gin.Context) {
	deviceId := c.Query("device_id")
	user, err := handler.uc.GetUserById(deviceId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
