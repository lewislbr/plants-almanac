package plant

import (
	"errors"
)

// ErrMissingData defines an error to use when there are missing required fields.
var ErrMissingData = errors.New("missing data")

// ErrNotFound defines an error to use when a plant is not found.
var ErrNotFound = errors.New("plant not found")
