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
		gs Generater
		r  Finder
	}
)

func NewService(gs Generater, r Finder) *service {
	return &service{gs, r}
}

func (s *service) Authenticate(cred user.Credentials) (string, error) {
	if cred.Email == "" || cred.Password == "" {
		return "", user.ErrMissingData
	}

	existUser, err := s.r.FindOne(cred.Email)
	if err != nil {
		return "", user.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Hash), []byte(cred.Password))
	if err != nil {
		return "", user.ErrInvalidPassword
	}

	token, err := s.gs.GenerateToken(existUser.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
