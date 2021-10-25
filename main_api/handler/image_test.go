package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/testutils"
)

type UploadImageForm struct {
	AlbumId  string
	ImageNum string
}

func TestUploadImages(t *testing.T) {
	tests := []struct {
		name                        string
		textForm                    UploadImageForm
		fakeStoreImages             func(images []*domain.Image) ([]int, error)
		fakeGetCoordinatesByAlbumId func(albumId int) ([]*domain.Coordinate, error)
		fakeS3Uploader              func(images []multipart.File, names []string) ([]string, error)
		want                        gin.H
		code                        int
		err                         bool
	}{
		{
			name: "success",
			textForm: UploadImageForm{
				AlbumId:  "1",
				ImageNum: "0",
			},
			fakeStoreImages: func(images []*domain.Image) ([]int, error) {
				return nil, nil
			},
			fakeGetCoordinatesByAlbumId: func(albumId int) ([]*domain.Coordinate, error) {
				return nil, nil
			},
			fakeS3Uploader: func(images []multipart.File, names []string) ([]string, error) {
				return nil, nil
			},
			want: gin.H{"data": "data"},
			code: http.StatusOK,
			err:  false,
		},
		{
			name: "album_id not found",
			textForm: UploadImageForm{
				AlbumId:  "",
				ImageNum: "0",
			},
			fakeStoreImages: func(images []*domain.Image) ([]int, error) {
				return nil, nil
			},
			fakeGetCoordinatesByAlbumId: func(albumId int) ([]*domain.Coordinate, error) {
				return nil, nil
			},
			fakeS3Uploader: func(images []multipart.File, names []string) ([]string, error) {
				return nil, nil
			},
			want: gin.H{"err": "`album_id` field not found"},
			code: http.StatusBadRequest,
			err:  true,
		},
		{
			name: "image_num not found",
			textForm: UploadImageForm{
				AlbumId:  "1",
				ImageNum: "",
			},
			fakeStoreImages: func(images []*domain.Image) ([]int, error) {
				return nil, nil
			},
			fakeGetCoordinatesByAlbumId: func(albumId int) ([]*domain.Coordinate, error) {
				return nil, nil
			},
			fakeS3Uploader: func(images []multipart.File, names []string) ([]string, error) {
				return nil, nil
			},
			want: gin.H{"err": "`image_num` field not found"},
			code: http.StatusBadRequest,
			err:  true,
		},
	}

	for _, tt := range tests {
		imageRepo := testutils.FakeImageRepository{
			FakeStoreImages: tt.fakeStoreImages,
		}
		coordinateRepo := testutils.FakeCoordinateRepository{
			FakeGetCoordinatesByAlbumId: tt.fakeGetCoordinatesByAlbumId,
		}
		s3service := testutils.FakeS3Service{
			FakeS3Uploader: tt.fakeS3Uploader,
		}

		imageHandler := NewImageHandler(imageRepo, s3service, coordinateRepo)

		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer
			w := multipart.NewWriter(&b)
			fw, _ := w.CreateFormField("album_id")
			_, _ = io.Copy(fw, strings.NewReader(tt.textForm.AlbumId))
			fw, _ = w.CreateFormField("image_num")
			_, _ = io.Copy(fw, strings.NewReader(tt.textForm.ImageNum))
			w.Close()

			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodPost,
				"/user",
				&b,
			)
			c.Request.Header.Set("Content-Type", w.FormDataContentType())
			imageHandler.UploadImages(c)

			var responseBody map[string]interface{}
			_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

			assert.Equal(t, tt.code, response.Code)
			if tt.err {
				assert.Equal(t, tt.want["err"], responseBody["err"])
			} else {
				assert.Equal(t, tt.want["data"], responseBody["data"])
			}
		})
	}
}
