package domain

import (
	"time"

	"github.com/jphacks/D_2106_2/utils"
)

type Album struct {
	Id                int       `json:"id"`
	UserId            string    `json:"userId"`
	Title             string    `json:"title"`
	StartedAt         time.Time `json:"starteAt"`
	EndedAt           time.Time `json:"endedAt"`
	IsPublic          bool      `json:"isPubliuc"`
	Spot              string    `json:"spot"`
	ThumbnailImageUrl string    `json:"thumbnailImageUrl"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (album *Album) ToResponse() *AlbumResponse {
	return &AlbumResponse{
		Id:                album.Id,
		UserId:            album.UserId,
		Title:             album.Title,
		StartedAt:         utils.TimeToUnix(album.StartedAt),
		EndedAt:           utils.TimeToUnix(album.EndedAt),
		IsPublic:          album.IsPublic,
		Spot:              album.Spot,
		ThumbnailImageUrl: album.ThumbnailImageUrl,
		CreatedAt:         utils.TimeToUnix(album.CreatedAt),
	}
}

type AlbumDB struct {
	Id               int
	UserId           string
	Title            string
	StartedAt        time.Time
	EndedAt          time.Time
	IsPublic         bool
	Spot             string
	ThumbnailImageId int
	CreatedAt        time.Time
}

func (AlbumDB) TableName() string {
	return "albums"
}

type AlbumResponse struct {
	Id                int    `json:"id"`
	UserId            string `json:"userId"`
	Title             string `json:"title"`
	StartedAt         int64  `json:"starteAt"`
	EndedAt           int64  `json:"endedAt"`
	IsPublic          bool   `json:"isPubliuc"`
	Spot              string `json:"spot"`
	ThumbnailImageUrl string `json:"thumbnailImage_url"`
	CreatedAt         int64  `json:"createdAt"`
}
