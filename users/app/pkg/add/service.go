package add

import (
	"users/pkg/entity"

	"github.com/google/uuid"
)

// Service provides user add operations
type Service interface {
	AddUser(entity.User)
}

// Repository provides access to the user storage
type Repository interface {
	InsertOne(entity.User)
}

type service struct {
	r Repository
}

// NewService creates an add service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddUser adds a new user
func (s *service) AddUser(user entity.User) {
	user.ID = uuid.New().String()

	s.r.InsertOne(user)
}
