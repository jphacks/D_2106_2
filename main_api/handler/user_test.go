package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/testutils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func GetDummyAuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} { return nil },
		Authenticator: func(c *gin.Context) (interface{}, error) {
			return nil, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	return authMiddleware
}

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name           string
		userReq        RegisteruserReq
		fakeCreateUser func(user *domain.User) (int, error)
		want           gin.H
		code           int
	}{
		{
			name: "success",
			userReq: RegisteruserReq{
				Username: "name",
				Password: "password",
			},
			fakeCreateUser: func(user *domain.User) (int, error) {
				return 1, nil
			},
			want: gin.H{"userId": 1},
			code: http.StatusOK,
		},
		{
			name: "failed username is null",
			userReq: RegisteruserReq{
				Password: "password",
			},
			want: gin.H{"err": errors.New("username or password field is null")},
			code: 500,
		},
		{
			name: "failed password is null",
			userReq: RegisteruserReq{
				Username: "name",
			},
			want: gin.H{"err": errors.New("username or password field is null")},
			code: 500,
		},
		{
			name: "failed register new user 1",
			userReq: RegisteruserReq{
				Username: "name",
				Password: "password",
			},
			fakeCreateUser: func(user *domain.User) (int, error) {
				return -1, errors.New("failed register new user")
			},
			want: gin.H{"err": errors.New("failed register new user")},
			code: 500,
		},
		{
			name: "failed register new user 2",
			userReq: RegisteruserReq{
				Username: "name",
				Password: "password",
			},
			fakeCreateUser: func(user *domain.User) (int, error) {
				return -1, nil
			},
			want: gin.H{"err": errors.New("regisster usesr failed")},
			code: 500,
		},
	}

	for _, tt := range tests {
		userRepo := testutils.FakeUserRepository{
			FakeCreateUser: tt.fakeCreateUser,
		}

		userHandler := NewUserHandler(userRepo)

		// auth middleware
		authMiddleware := GetDummyAuthMiddleware()

		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.userReq)
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodPost,
				"/user",
				bytes.NewBuffer(body),
			)
			userHandler.RegisterUser(c, authMiddleware)

			var responseBody map[string]interface{}
			_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

			assert.Equal(t, tt.code, response.Code)
			assert.Equal(t, tt.want["usesrId"], responseBody["usesrId"])
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	tests := []struct {
		name            string
		fakeGetAllUsers func() ([]*domain.User, error)
		want            []*domain.User
		code            int
		isError         bool
		wantError       error
	}{
		{
			name: "success",
			fakeGetAllUsers: func() ([]*domain.User, error) {
				return []*domain.User{
					{
						Id:       1,
						Name:     "name",
						Password: "pass",
					},
				}, nil
			},
			want: []*domain.User{
				{
					Id:       1,
					Name:     "name",
					Password: "pass",
				},
			},
			code:    200,
			isError: false,
		},
		{
			name: "failed get all users 1",
			fakeGetAllUsers: func() ([]*domain.User, error) {
				return nil, errors.New("get all user error")
			},
			code:      500,
			isError:   true,
			wantError: errors.New("get all user error"),
		},
		{
			name: "failed get all users 2",
			fakeGetAllUsers: func() ([]*domain.User, error) {
				return []*domain.User{}, nil
			},
			code:      500,
			isError:   true,
			wantError: errors.New("users not found"),
		},
	}

	for _, tt := range tests {
		userRepo := testutils.FakeUserRepository{
			FakeGetAllUsers: tt.fakeGetAllUsers,
		}

		userHandler := NewUserHandler(userRepo)

		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodGet,
				"/users",
				nil,
			)
			userHandler.GetAllUsers(c)

			assert.Equal(t, tt.code, response.Code)
			if !tt.isError {
				var responseBody []*domain.User
				_ = json.Unmarshal(response.Body.Bytes(), &responseBody)
				assert.Equal(t, tt.want, responseBody)
			} else {
				var responseBody map[string]interface{}
				_ = json.Unmarshal(response.Body.Bytes(), &responseBody)
				assert.Equal(t, tt.wantError.Error(), responseBody["err"])
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name        string
		fakeGetUser func(userId int) (*domain.User, error)
		want        *domain.User
		code        int
		isError     bool
		wantError   error
	}{
		{
			name: "success",
			fakeGetUser: func(userId int) (*domain.User, error) {
				return &domain.User{
					Id:       1,
					Name:     "name",
					Password: "pass",
				}, nil
			},
			want: &domain.User{
				Id:       1,
				Name:     "name",
				Password: "pass",
			},
			code:    200,
			isError: false,
		},
		{
			name: "failed get user by id",
			fakeGetUser: func(userId int) (*domain.User, error) {
				return nil, errors.New("get user error")
			},
			code:      500,
			isError:   true,
			wantError: errors.New("get user error"),
		},
	}

	for _, tt := range tests {
		userRepo := testutils.FakeUserRepository{
			FakeGetUserById: tt.fakeGetUser,
		}

		userHandler := NewUserHandler(userRepo)

		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodGet,
				"/user",
				nil,
			)

			params := c.Request.URL.Query()
			params.Add("id", "1")
			c.Request.URL.RawQuery = params.Encode()

			userHandler.GetUser(c)

			assert.Equal(t, tt.code, response.Code)
			if !tt.isError {
				var responseBody *domain.User
				_ = json.Unmarshal(response.Body.Bytes(), &responseBody)
				assert.Equal(t, tt.want, responseBody)
			} else {
				var responseBody map[string]interface{}
				_ = json.Unmarshal(response.Body.Bytes(), &responseBody)
				assert.Equal(t, tt.wantError.Error(), responseBody["err"])
			}
		})
	}
}
