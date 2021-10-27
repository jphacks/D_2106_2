package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/testutils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name           string
		userReq        interface{}
		fakeCreateUser func(user *domain.User) (int, error)
		want           gin.H
		code           int
	}{
		{
			name: "success",
			userReq: RegisterUserRequest{
				Username: "test-name",
				DeviceId: "test-device-id",
			},
			fakeCreateUser: func(user *domain.User) (int, error) {
				return 1, nil
			},
			want: gin.H{"data": &RegisterUserRespose{UserId: 1}},
			code: http.StatusOK,
		},
		{
			name: "failed username is null",
			userReq: RegisterUserRequest{
				DeviceId: "test-device-id",
			},
			want: gin.H{"error": FieldIsNull.Error()},
			code: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		userRepo := testutils.FakeUserRepository{
			FakeCreateUser: tt.fakeCreateUser,
		}

		userHandler := NewUserHandler(userRepo)

		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.userReq)
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodPost,
				"/user",
				bytes.NewBuffer(body),
			)
			userHandler.RegisterUser(c)

			assert.Equal(t, tt.code, response.Code)
			wantJson, err := json.Marshal(tt.want)
			assert.NoError(t, err)
			assert.Equal(t, wantJson, response.Body.Bytes())
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name                  string
		fakeGetUserByDeviceId func(deviceId string) (*domain.User, error)
		want                  gin.H
		code                  int
	}{
		{
			name: "success",
			fakeGetUserByDeviceId: func(deviceId string) (*domain.User, error) {
				return &domain.User{Id: 1}, nil
			},
			want: gin.H{"data": &domain.User{Id: 1}},
			code: http.StatusOK,
		},
		{
			name: "failed get user by id",
			fakeGetUserByDeviceId: func(deviceId string) (*domain.User, error) {
				return nil, errors.New("get user error")
			},
			want: gin.H{"error": errors.New("get user error").Error()},
			code: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		userRepo := testutils.FakeUserRepository{
			FakeGetUserByDeviceId: tt.fakeGetUserByDeviceId,
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
			params.Add("device_id", "dummy-device-id")
			c.Request.URL.RawQuery = params.Encode()

			userHandler.GetUser(c)

			assert.Equal(t, tt.code, response.Code)
			wantJson, err := json.Marshal(tt.want)
			assert.NoError(t, err)
			assert.Equal(t, wantJson, response.Body.Bytes())
		})
	}
}
