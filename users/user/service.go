package user

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	postgresRepository interface {
		Insert(User) error
		CheckExists(string) (bool, error)
		Find(string) (User, error)
		GetInfo(string) (Info, error)
	}

	service struct {
		postgres postgresRepository
	}
)

func NewService(postgres postgresRepository) *service {
	return &service{postgres}
}

func (s *service) Create(new User) error {
	if new.Name == "" || new.Email == "" || new.Password == "" {
		return ErrMissingData
	}

	exists, err := s.postgres.CheckExists(new.Email)
	if err != nil {
		return err
	}
	if exists {
		return ErrUserExists
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()

	hash, err := bcrypt.GenerateFromPassword([]byte(new.Password), 10)
	if err != nil {
		return err
	}

	new.Hash = string(hash)

	return s.postgres.Insert(new)
}

func (s *service) Authenticate(cred Credentials) (string, error) {
	if cred.Email == "" || cred.Password == "" {
		return "", ErrMissingData
	}

	existUser, err := s.postgres.Find(cred.Email)
	if err != nil {
		return "", ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Hash), []byte(cred.Password))
	if err != nil {
		return "", ErrInvalidPassword
	}

	return existUser.ID, nil
}

func (s *service) Info(userID string) (Info, error) {
	if userID == "" {
		return Info{}, ErrMissingData
	}

	userInfo, err := s.postgres.GetInfo(userID)
	if err != nil {
		return Info{}, ErrNotFound
	}

	return userInfo, nil
}
