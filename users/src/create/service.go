package create

import (
	"time"

	u "users/src/user"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Service defines a service to create a user.
type Service interface {
	Create(u.User) error
}

type repository interface {
	FindOne(string) (u.User, error)
	InsertOne(u.User) (interface{}, error)
}

type service struct {
	r repository
}

// NewService creates a create service with the necessary dependencies.
func NewService(r repository) Service {
	return &service{r}
}

// Create creates a new user.
func (s *service) Create(new u.User) error {
	_, err := s.r.FindOne(new.Email)
	if err == nil {
		return u.ErrUserExists
	}

	new.ID = u.ID(uuid.New().String())
	new.CreatedAt = time.Now().UTC()

	hash, err := bcrypt.GenerateFromPassword([]byte(new.Password), 10)
	if err != nil {
		return errors.Wrap(err, "")
	}

	new.Hash = string(hash)

	_, err = s.r.InsertOne(new)
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}
