package user

import (
	"errors"
)

var ErrMissingData = errors.New("missing data")

var ErrNotFound = errors.New("user not found")

var ErrUserExists = errors.New("email already registered")

var ErrInvalidPassword = errors.New("invalid password")

var ErrInvalidToken = errors.New("invalid token")
