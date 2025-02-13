package request

import "errors"

var (
	// SongNotFound is an error type for when the song api returns
	// `lyrics not found`
	SongNotFound = errors.New("song not found")
)
