package user

import "time"

type (
	// Ideally the JSON and BSON tags should be defined in an specific entity
	// for the server and storage components, respectively, but this being a small
	// service they are defined here for simplicity.
	User struct {
		ID        string    `json:"id" bson:"_id"`
		Name      string    `json:"name" bson:"name"`
		Email     string    `json:"email" bson:"email"`
		Password  string    `json:"password" bson:"-"`
		Hash      string    `json:"-" bson:"hash"`
		CreatedAt time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	}
	Credentials struct {
		Email    string
		Password string
	}
)
