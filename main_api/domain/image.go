package domain

import "time"

type Image struct {
	Id           int
	Url          string
	Name         string
	AlbumId      int
	CreatedAt    time.Time
	CoordinateId int
}
