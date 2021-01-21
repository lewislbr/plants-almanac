package user

import (
	"errors"
	"time"
)

// User defines the properties of a user.
// Ideally the JSON and BSON tags should be defined in an specific entity
// for the API and storage components, respectively, but this being a small
// service they are defined here for simplicity.
type User struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"-"`
	Hash      string    `json:"-" bson:"hash"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// Credentials defines the data needed to authenticate.
type Credentials struct {
	Email    string
	Password string
}

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
