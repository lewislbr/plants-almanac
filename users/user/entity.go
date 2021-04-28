package user

import "time"

type (
	// Ideally the JSON and BSON tags should be defined in an specific entity
	// for the server and storage components, respectively, but this being a small
	// service they are defined here for simplicity.
	User struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Hash      string    `json:"-"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	Credentials struct {
		Email    string
		Password string
	}
	Info struct {
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
	}
)
