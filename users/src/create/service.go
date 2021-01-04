package create

import (
	"time"

	u "users/src/user"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Service provides user creation operations
type Service interface {
	Create(u.User) error
}

// Repository provides access to the user storage
type Repository interface {
	FindOne(string) (u.User, error)
	InsertOne(u.User) (interface{}, error)
}

type service struct {
	r Repository
}

// Create creates a new user
func (s *service) Create(newUser u.User) error {
	_, err := s.r.FindOne(newUser.Email)
	if err == nil {
		return u.ErrUserExists
	}

	newUser.ID = u.ID(uuid.New().String())
	newUser.CreatedAt = time.Now().UTC()

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		return errors.Wrap(err, "")
	}

	newUser.Hash = string(hash)

	_, err = s.r.InsertOne(newUser)
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}

// NewService creates a create service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
