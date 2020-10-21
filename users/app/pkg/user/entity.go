package user

import "errors"

// User defines the properties of a user
type User struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}

// ID defines the id of a user
type ID string

// ErrNotFound creates an error to use when a user is not found
var ErrNotFound = errors.New("user not found")
