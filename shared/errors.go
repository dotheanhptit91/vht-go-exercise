package shared

import "errors"

var (
	ErrDataNotFound   = errors.New("data not found")
	ErrNoDataToUpdate = errors.New("no data to update")
)