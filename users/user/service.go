package user

import (
	"fmt"
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
		return fmt.Errorf("error creating user: %w", ErrMissingData)
	}

	exists, err := s.postgres.CheckExists(new.Email)
	if err != nil {
		return fmt.Errorf("error checking user: %w", ErrMissingData)
	}
	if exists {
		return fmt.Errorf("error creating user: %w", ErrUserExists)
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()

	hash, err := bcrypt.GenerateFromPassword([]byte(new.Password), 10)
	if err != nil {
		return fmt.Errorf("error generating password: %w", err)
	}

	new.Hash = string(hash)

	err = s.postgres.Insert(new)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}

func (s *service) Authenticate(cred Credentials) (string, error) {
	if cred.Email == "" || cred.Password == "" {
		return "", fmt.Errorf("error authenticating user: %w", ErrMissingData)

	}

	existUser, err := s.postgres.Find(cred.Email)
	if err != nil {
		return "", fmt.Errorf("error finding user: %w", ErrNotFound)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Hash), []byte(cred.Password))
	if err != nil {
		return "", fmt.Errorf("error validating password: %w", ErrInvalidPassword)
	}

	return existUser.ID, nil
}

func (s *service) Info(userID string) (Info, error) {
	if userID == "" {
		return Info{}, fmt.Errorf("error getting user info: %w", ErrMissingData)
	}

	userInfo, err := s.postgres.GetInfo(userID)
	if err != nil {
		return Info{}, fmt.Errorf("error getting user info: %w", ErrNotFound)
	}

	return userInfo, nil
}
