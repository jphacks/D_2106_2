package handler

import "errors"

var (
	InvalidRequest               = errors.New("Invalid request parameter")
	InvalidCoordinate            = errors.New("Invalid coordinates")
	UserFieldIsNull              = errors.New("Username or deviceId field is null")
	FailedRegisterUser           = errors.New("Register usesr failed")
	FailedGetAlbum               = errors.New("Failed getting albums")
	FailedCreateNewAlbum         = errors.New("Failed create new album")
	FailedClustering             = errors.New("Failed clustering")
	FailedUpdateThumbnailAndSpot = errors.New("Failed to update thumbnail and spot")
)
