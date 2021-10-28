package domain

import (
	"time"
)

type Album struct {
	Id                int
	UserId            string
	Title             string
	StartedAt         time.Time
	EndedAt           time.Time
	IsPublic          bool
	ThumbnailImageUrl string
}

type AlbumDB struct {
	Id               int
	UserId           string
	Title            string
	StartedAt        time.Time
	EndedAt          time.Time
	IsPublic         bool
	ThumbnailImageId int
}
