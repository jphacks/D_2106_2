package domain

import (
	"time"
)

type Album struct {
	Id               int
	UserId           int
	Title            string
	StartedAt        time.Time
	EndedAt          time.Time
	IsPublic         bool
	ThumbnailImageId int
}
