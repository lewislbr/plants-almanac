package user

import (
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

// CreateService defines a service to create a user.
type CreateService interface {
	Create(User) error
}

// AuthenticateService defines a service to authenticate a user.
type AuthenticateService interface {
	Authenticate(cred Credentials) (string, error)
}

// AuthorizeService defines a service to authorize a user.
type AuthorizeService interface {
	Authorize(string) (string, error)
}

// Repository defines storage operations.
type Repository interface {
	InsertOne(User) (interface{}, error)
	FindOne(string) (User, error)
}
