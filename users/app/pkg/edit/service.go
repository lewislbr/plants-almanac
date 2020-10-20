package edit

import (
	"users/pkg/entity"
	"errors"
)

var errNotFound = errors.New("user not found")

// Service provides user edit operations
type Service interface {
	EditUser(string, entity.User) error
}

// Repository provides access to the user storage
type Repository interface {
	EditOne(string, entity.User)
	FindOne(string) (*entity.User, bool)
}

type service struct {
	r Repository
}

// NewService creates an edit service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// EditUser edits an user
func (s *service) EditUser(id string, user entity.User) error {
	existingUser, ok := s.r.FindOne(id)
	if !ok {
		return errNotFound
	}
	if user.Name == "" {
		user.Name = existingUser.Name
	}

	user = entity.User{ID: id, Name: user.Name}

	s.r.EditOne(id, user)

	return nil
}
