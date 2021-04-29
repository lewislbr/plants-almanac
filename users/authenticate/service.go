package authenticate

import (
	"users/user"

	"golang.org/x/crypto/bcrypt"
)

type (
	repository interface {
		FindOne(string) (user.User, error)
	}

	service struct {
		repo repository
	}
)

func NewService(repo repository) *service {
	return &service{repo}
}

func (s *service) Authenticate(cred user.Credentials) (string, error) {
	if cred.Email == "" || cred.Password == "" {
		return "", user.ErrMissingData
	}

	existUser, err := s.repo.FindOne(cred.Email)
	if err != nil {
		return "", user.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Hash), []byte(cred.Password))
	if err != nil {
		return "", user.ErrInvalidPassword
	}

	return existUser.ID, nil
}
