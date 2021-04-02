package authenticate

import (
	"users/storage"
	"users/user"

	"golang.org/x/crypto/bcrypt"
)

type (
	Generater interface {
		GenerateJWT(string) (string, error)
	}

	service struct {
		gs Generater
		r  storage.Repository
	}
)

func NewService(gs Generater, r storage.Repository) *service {
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

	jwt, err := s.gs.GenerateJWT(existUser.ID)
	if err != nil {
		return "", err
	}

	return jwt, nil
}
