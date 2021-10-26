package domain

import (
	"time"
)

type Coordinate struct {
	Id        int
	AlbumId   int
	Timestamp time.Time
	Latitude  float64
	Longitude float64
	IsShow    bool
}

type Location struct {
	Latitude  float64
	Longitude float64
	Timestamp int64
}
