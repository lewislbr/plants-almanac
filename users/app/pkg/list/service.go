package list

import (
	"users/pkg/entity"
	"errors"
)

var errNotFound = errors.New("user not found")

// Service provides user list operations
type Service interface {
	GetUser(string) (*entity.User, error)
}

// Repository provides access to the user storage
type Repository interface {
	FindOne(string) (*entity.User, bool)
}

type service struct {
	r Repository
}

// NewService creates a list service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetUser returns an user
func (s *service) GetUser(id string) (*entity.User, error) {
	user, ok := s.r.FindOne(id)
	if !ok {
		return nil, errNotFound
	}

	return user, nil
}
