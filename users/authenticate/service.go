package authenticate

import (
	"users/generate"
	"users/storage"

	"users/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthenticateService interface {
	Authenticate(cred user.Credentials) (string, error)
}

type authenticateService struct {
	gs generate.GenerateService
	r  storage.Repository
}

func NewAuthenticateService(gs generate.GenerateService, r storage.Repository) *authenticateService {
	return &authenticateService{gs, r}
}

func (ns *authenticateService) Authenticate(cred user.Credentials) (string, error) {
	if cred.Email == "" || cred.Password == "" {
		return "", user.ErrMissingData
	}

	existUser, err := ns.r.FindOne(cred.Email)
	if err != nil {
		return "", user.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Hash), []byte(cred.Password))
	if err != nil {
		return "", user.ErrInvalidPassword
	}

	jwt, err := ns.gs.GenerateJWT(existUser.ID)
	if err != nil {
		return "", err
	}

	return jwt, nil
}
