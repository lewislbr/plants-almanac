package list

import (
	u "users/pkg/user"
)

// Service provides user list operations
type Service interface {
	ListUser(u.ID) (*u.User, error)
}

// Repository provides access to the user storage
type Repository interface {
	FindOne(u.ID) (*u.User, bool)
}

type service struct {
	r Repository
}

// ListUser returns a user
func (s *service) ListUser(id u.ID) (*u.User, error) {
	user, ok := s.r.FindOne(id)
	if !ok {
		return nil, u.ErrNotFound
	}

	return user, nil
}

// NewService creates a list service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
