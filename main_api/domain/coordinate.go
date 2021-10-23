package domain

type Coordinate struct {
	Id        int
	AlbumId   int
	Timestamp string
	Latitude  float32
	Longitude float32
}

type Location struct {
	Latitude  float32
	Longitude float32
	Timestamp string
}
