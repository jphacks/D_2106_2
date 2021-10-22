package handler

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
	"github.com/jphacks/D_2106_2/usecase"
)

var identityKey = "id"

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type AuthHandler struct {
	uc usecase.UserUsecase
}

func NewAuthHandler(userRepo repository.UserRepository) *AuthHandler {
	uc := usecase.UserUsecase{UserRepo: userRepo}

	return &AuthHandler{uc: uc}
}

func (handler *AuthHandler) IdentityHandler(c *gin.Context) *domain.User {
	claims := jwt.ExtractClaims(c)
	return &domain.User{
		Name: claims[identityKey].(string),
	}
}

func (handler *AuthHandler) AuthenticateHandler(c *gin.Context) (*domain.User, error) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password

	user, err := handler.uc.Login(userID, password)
	if err != nil {
		log.Print(err)
		return nil, jwt.ErrFailedAuthentication
	}

	return user, nil
}

func (handler *AuthHandler) AuthorizeHandler(c *gin.Context, data interface{}) bool {
	v, ok := data.(*domain.User)
	if !ok {
		return false
	}

	isExists, err := handler.uc.CheckUserExist(v.Name)
	if err != nil {
		return false
	}

	return isExists
}

func (handler *AuthHandler) UnauthorizeHandler(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
