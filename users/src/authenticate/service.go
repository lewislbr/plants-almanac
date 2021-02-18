package authenticate

import (
	u "users/src/user"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type authenticateService struct {
	g u.GenerateService
	r u.Repository
}

// NewAuthenticateService initializes an authentication service with the necessary dependencies.
func NewAuthenticateService(g u.GenerateService, r u.Repository) u.AuthenticateService {
	return authenticateService{g, r}
}

// Authenticate authenticates a user and issues a JWT.
func (s authenticateService) Authenticate(cred u.Credentials) (string, error) {
	if cred.Email == "" || cred.Password == "" {
		return "", u.ErrMissingData
	}

	existUser, err := s.r.FindOne(cred.Email)
	if err != nil {
		return "", u.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Hash), []byte(cred.Password))
	if err != nil {
		return "", u.ErrInvalidPassword
	}

	jwt, err := s.g.GenerateJWT(existUser.ID)
	if err != nil {

		return "", errors.Wrap(err, "")
	}

	return jwt, nil
}
