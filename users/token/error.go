package token

import "errors"

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrMissingData  = errors.New("missing data")
)
