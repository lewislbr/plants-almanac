package user

import "time"

type (
	User struct {
		ID        string    `firestore:"id"`
		Name      string    `firestore:"name"`
		Email     string    `firestore:"email"`
		Hash      string    `firestore:"hash"`
		CreatedAt time.Time `firestore:"created_at"`
	}
	New struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	Credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	Info struct {
		Name      string    `json:"name" firestore:"name"`
		Email     string    `json:"email" firestore:"email"`
		CreatedAt time.Time `json:"created_at" firestore:"created_at"`
	}
)
