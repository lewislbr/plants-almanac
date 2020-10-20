package delete

import (
	"users/pkg/entity"
	"errors"
)

var errNotFound = errors.New("user not found")

// Service provides user delete operations
type Service interface {
	DeleteUser(string) error
}

// Repository provides access to the user storage
type Repository interface {
	DeleteOne(string)
	FindOne(string) (*entity.User, bool)
}

type service struct {
	r Repository
}

// NewService creates a delete service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// DeleteUser deletes an user
func (s *service) DeleteUser(id string) error {
	_, ok := s.r.FindOne(id)
	if !ok {
		return errNotFound
	}

	s.r.DeleteOne(id)

	return nil
}
