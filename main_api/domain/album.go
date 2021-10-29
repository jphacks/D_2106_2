package domain

import (
	"time"
)

type Album struct {
	Id                int       `json:"id"`
	UserId            string    `json:"userId"`
	Title             string    `json:"title"`
	StartedAt         time.Time `json:"starteAt"`
	EndedAt           time.Time `json:"endedAt"`
	IsPublic          bool      `json:"isPubliuc"`
	ThumbnailImageUrl string    `json:"thumbnailImage_url"`
	CreatedAt         time.Time `json:"createdAt"`
}

type AlbumDB struct {
	Id               int
	UserId           string
	Title            string
	StartedAt        time.Time
	EndedAt          time.Time
	IsPublic         bool
	ThumbnailImageId int
	CreatedAt        time.Time
}

func (AlbumDB) TableName() string {
	return "albums"
}
