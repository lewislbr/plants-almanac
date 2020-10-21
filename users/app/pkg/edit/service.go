package edit

import (
	u "users/pkg/user"
)

// Service provides user edit operations
type Service interface {
	EditUser(u.ID, u.User) error
}

// Repository provides access to the user storage
type Repository interface {
	EditOne(u.ID, u.User)
	FindOne(u.ID) (*u.User, bool)
}

type service struct {
	r Repository
}

// EditUser edits a user
func (s *service) EditUser(id u.ID, user u.User) error {
	existingUser, ok := s.r.FindOne(id)
	if !ok {
		return u.ErrNotFound
	}
	if user.Name == "" {
		user.Name = existingUser.Name
	}

	user = u.User{ID: id, Name: user.Name}

	s.r.EditOne(id, user)

	return nil
}

// NewService creates an edit service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
