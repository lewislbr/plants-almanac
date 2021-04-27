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
		repo InserterChecker
	}
)

func NewService(repo InserterChecker) *service {
	return &service{repo}
}

func (s *service) Create(new user.User) error {
	if new.Name == "" || new.Email == "" || new.Password == "" {
		return user.ErrMissingData
	}

	exists, err := s.repo.CheckExists(new.Email)
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

	return s.repo.InsertOne(new)
}
