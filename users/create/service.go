package create

import (
	"time"

	"users/storage"
	"users/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	r storage.Repository
}

func NewService(r storage.Repository) *service {
	return &service{r}
}

func (s *service) Create(new user.User) error {
	if new.Name == "" || new.Email == "" || new.Password == "" {
		return user.ErrMissingData
	}

	_, err := s.r.FindOne(new.Email)
	if err == nil {
		return user.ErrUserExists
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()

	hash, err := bcrypt.GenerateFromPassword([]byte(new.Password), 10)
	if err != nil {
		return err
	}

	new.Hash = string(hash)

	_, err = s.r.InsertOne(new)
	if err != nil {
		return err
	}

	return nil
}
