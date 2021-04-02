package plant

import "errors"

var (
	ErrMissingData = errors.New("missing data")
	ErrNotFound    = errors.New("plant not found")
)
