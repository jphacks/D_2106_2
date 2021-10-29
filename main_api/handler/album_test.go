package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/testutils"
	"github.com/stretchr/testify/assert"
)

func TestGetAllAlbums(t *testing.T) {
	tests := []struct {
		name             string
		fakeGetAllAlbums func() ([]*domain.Album, error)
		want             gin.H
		code             int
	}{
		{
			name: "success",
			fakeGetAllAlbums: func() ([]*domain.Album, error) {
				return []*domain.Album{
					{Id: 1},
					{Id: 2},
					{Id: 3},
				}, nil
			},
			want: gin.H{"data": GetAllAlbumsResponse{
				[]*domain.AlbumResponse{
					{Id: 1},
					{Id: 2},
					{Id: 3},
				},
			}},
			code: http.StatusOK,
		},
	}

	for _, tt := range tests {
		albumRepo := testutils.FakeAlbumRepository{
			FakeGetAllAlbums: tt.fakeGetAllAlbums,
			FakeGetAlbumsByUsers: func(userId string) ([]*domain.Album, error) {
				return nil, nil
			},
		}
		coordinateRepo := testutils.FakeCoordinateRepository{}
		imageRepo := testutils.FakeImageRepository{}

		albumHandler := NewAlbumHandler(albumRepo, coordinateRepo, imageRepo)

		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodGet,
				"/albuls",
				nil,
			)

			albumHandler.GetAllAlbums(c)

			assert.Equal(t, tt.code, response.Code)
			// wantJson, err := json.Marshal(tt.want)
			// assert.NoError(t, err)
			// assert.Equal(t, wantJson, response.Body.Bytes())
		})
	}
}
