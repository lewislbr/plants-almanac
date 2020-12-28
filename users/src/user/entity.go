package user

import (
	"errors"
	"time"
)

// User defines the properties of a user
type User struct {
	ID        ID        `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"-"`
	Hash      string    `json:"-" bson:"hash"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// ID defines the id of a user
type ID string

// Credentials defines the data needed to authenticate
type Credentials struct {
	Email    string
	Password string
}

// ErrUserExists creates an error to use when a user already exists when
// creating a user
var ErrUserExists = errors.New("user already exists")

// ErrInvalidPassword creates an error to use when the user password does
// not match
var ErrInvalidPassword = errors.New("invalid password")

// ErrMissingData creates an error to use when fields required to create a
// new user are missing
var ErrMissingData = errors.New("missing data")

// ErrNotFound creates an error to use when a user is not found
var ErrNotFound = errors.New("user not found")
