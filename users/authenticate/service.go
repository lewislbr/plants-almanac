package authenticate

import (
	u "users/user"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type authenticateService struct {
	gs u.GenerateService
	r  u.Repository
}

// NewAuthenticateService initializes an authentication service with the necessary dependencies.
func NewAuthenticateService(gs u.GenerateService, r u.Repository) *authenticateService {
	return &authenticateService{gs, r}
}

// Authenticate authenticates a user and issues a JWT.
func (ns *authenticateService) Authenticate(cred u.Credentials) (string, error) {
	if cred.Email == "" || cred.Password == "" {
		return "", u.ErrMissingData
	}

	existUser, err := ns.r.FindOne(cred.Email)
	if err != nil {
		return "", u.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Hash), []byte(cred.Password))
	if err != nil {
		return "", u.ErrInvalidPassword
	}

	jwt, err := ns.gs.GenerateJWT(existUser.ID)
	if err != nil {

		return "", errors.Wrap(err, "")
	}

	return jwt, nil
}
