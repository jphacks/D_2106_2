package domain

import "time"

type Image struct {
	Id           int
	Name         string
	Url          string
	AlbumId      int
	CreatedAt    time.Time
	CoordinateId int
}
