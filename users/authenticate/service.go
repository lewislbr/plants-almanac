package authenticate

import (
	"users/user"

	"golang.org/x/crypto/bcrypt"
)

type (
	Generater interface {
		GenerateToken(string) (string, error)
	}

	Finder interface {
		FindOne(string) (user.User, error)
	}

	service struct {
		svc  Generater
		repo Finder
	}
)

func NewService(svc Generater, repo Finder) *service {
	return &service{svc, repo}
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

	token, err := s.svc.GenerateToken(existUser.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
