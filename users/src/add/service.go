package add

import (
	u "users/src/user"

	"github.com/google/uuid"
)

// Service provides user add operations
type Service interface {
	AddUser(u.User)
}

// Repository provides access to the user storage
type Repository interface {
	InsertOne(u.User)
}

type service struct {
	r Repository
}

// AddUser adds a user
func (s *service) AddUser(user u.User) {
	user.ID = u.ID(uuid.New().String())

	s.r.InsertOne(user)
}

// NewService creates an add service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
