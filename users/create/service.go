package create

import (
	"time"

	"users/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	InserterChecker interface {
		InsertOne(user.User) error
		CheckExists(string) (bool, error)
	}

	service struct {
		r InserterChecker
	}
)

func NewService(r InserterChecker) *service {
	return &service{r}
}

func (s *service) Create(new user.User) error {
	if new.Name == "" || new.Email == "" || new.Password == "" {
		return user.ErrMissingData
	}

	exists, err := s.r.CheckExists(new.Email)
	if err != nil {
		return err
	}
	if exists {
		return user.ErrUserExists
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()

	hash, err := bcrypt.GenerateFromPassword([]byte(new.Password), 10)
	if err != nil {
		return err
	}

	new.Hash = string(hash)

	err = s.r.InsertOne(new)
	if err != nil {
		return err
	}

	return nil
}
