package user

import (
	"errors"
)

var (
	ErrMissingData     = errors.New("missing data")
	ErrNotFound        = errors.New("user not found")
	ErrUserExists      = errors.New("email already registered")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidToken    = errors.New("invalid token")
)
