package user

import (
	"errors"
)

// ErrMissingData defines an error to use when there are missing required fields.
var ErrMissingData = errors.New("missing data")

// ErrNotFound defines an error to use when a user is not found.
var ErrNotFound = errors.New("user not found")

// ErrUserExists defines an error to use when a user already exists.
var ErrUserExists = errors.New("email already registered")

// ErrInvalidPassword defines an error to use when the user password does
// not match.
var ErrInvalidPassword = errors.New("invalid password")

// ErrInvalidToken defines an error to use when the user token is not
// valid.
var ErrInvalidToken = errors.New("invalid token")
