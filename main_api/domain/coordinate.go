package domain

type Coordinate struct {
	Id        int
	AlbumId   int
	Timestamp string
	Latitude  float64
	Longitude float64
	IsShow bool
}

type Location struct {
	Latitude  float64
	Longitude float64
	Timestamp string
}
