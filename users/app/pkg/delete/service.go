package delete

import (
	u "users/pkg/user"
)

// Service provides user delete operations
type Service interface {
	DeleteUser(u.ID) error
}

// Repository provides access to the user storage
type Repository interface {
	DeleteOne(u.ID)
	FindOne(u.ID) (*u.User, bool)
}

type service struct {
	r Repository
}

// DeleteUser deletes a user
func (s *service) DeleteUser(id u.ID) error {
	_, ok := s.r.FindOne(id)
	if !ok {
		return u.ErrNotFound
	}

	s.r.DeleteOne(id)

	return nil
}

// NewService creates a delete service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
