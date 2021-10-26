package handler

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/jphacks/D_2106_2/domain"
// 	"github.com/jphacks/D_2106_2/testutils"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/assert/v2"
// )

// func TestGetAllAlbums(t *testing.T) {
// 	tests := []struct {
// 		name             string
// 		fakeGetAllAlbums func() ([]*domain.Album, error)
// 		code             int
// 		want             gin.H
// 		isError          bool
// 	}{
// 		{
// 			name: "success",
// 			fakeGetAllAlbums: func() ([]*domain.Album, error) {
// 				return []*domain.Album{{Id: 1}}, nil
// 			},
// 			want:    gin.H{"data": []*domain.Album{{Id: 1}}},
// 			code:    200,
// 			isError: false,
// 		},
// 		{
// 			name: "failed get all users 1",
// 			fakeGetAllAlbums: func() ([]*domain.Album, error) {
// 				return nil, nil
// 			},
// 			want:    gin.H{"err": ""},
// 			code:    500,
// 			isError: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		albumRepo := testutils.FakeAlbumRepository{
// 			FakeGetAllAlbums: tt.fakeGetAllAlbums,
// 		}

// 		userHandler := NewAlbumHandler(albumRepo)

// 		t.Run(tt.name, func(t *testing.T) {
// 			response := httptest.NewRecorder()
// 			c, _ := gin.CreateTestContext(response)
// 			c.Request, _ = http.NewRequest(
// 				http.MethodGet,
// 				"/users",
// 				nil,
// 			)
// 			userHandler.GetAllUsers(c)

// 			assert.Equal(t, tt.code, response.Code)
// 			if !tt.isError {
// 				var responseBody []*domain.User
// 				_ = json.Unmarshal(response.Body.Bytes(), &responseBody)
// 				assert.Equal(t, tt.want, responseBody)
// 			} else {
// 				var responseBody map[string]interface{}
// 				_ = json.Unmarshal(response.Body.Bytes(), &responseBody)
// 				assert.Equal(t, tt.wantError.Error(), responseBody["err"])
// 			}
// 		})
// 	}
// }
