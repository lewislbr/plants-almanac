package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	userRepo interface {
		Insert(User) error
		CheckExists(string) (bool, error)
		Find(string) (User, error)
		GetInfo(string) (Info, error)
	}

	userService struct {
		userRepo userRepo
	}
)

func NewService(userRepo userRepo) *userService {
	return &userService{userRepo}
}

func (s *userService) Create(new New) error {
	if new.Name == "" || new.Email == "" || new.Password == "" {
		return fmt.Errorf("error creating user: %w", ErrMissingData)
	}

	exists, err := s.userRepo.CheckExists(new.Email)
	if err != nil {
		return fmt.Errorf("error checking user: %w", err)
	}
	if exists {
		return fmt.Errorf("error creating user: %w", ErrUserExists)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(new.Password), 10)
	if err != nil {
		return fmt.Errorf("error generating password: %w", err)
	}

	user := User{
		ID:        uuid.New().String(),
		Name:      new.Name,
		Email:     new.Email,
		Hash:      string(hash),
		CreatedAt: time.Now().UTC(),
	}

	err = s.userRepo.Insert(user)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}

func (s *userService) Authenticate(cred Credentials) (string, error) {
	if cred.Email == "" || cred.Password == "" {
		return "", fmt.Errorf("error authenticating user: %w", ErrMissingData)

	}

	existUser, err := s.userRepo.Find(cred.Email)
	if err != nil {
		return "", fmt.Errorf("error finding user: %w", err)
	}
	if existUser == (User{}) {
		return "", fmt.Errorf("error finding user: %w", ErrNotFound)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Hash), []byte(cred.Password))
	if err != nil {
		return "", fmt.Errorf("error validating password: %w", ErrInvalidPassword)
	}

	return existUser.ID, nil
}

func (s *userService) Info(userID string) (Info, error) {
	if userID == "" {
		return Info{}, fmt.Errorf("error getting user info: %w", ErrMissingData)
	}

	userInfo, err := s.userRepo.GetInfo(userID)
	if err != nil {
		return Info{}, fmt.Errorf("error getting user info: %w", err)
	}
	if userInfo == (Info{}) {
		return Info{}, fmt.Errorf("error getting user info: %w", ErrNotFound)
	}

	return userInfo, nil
}
