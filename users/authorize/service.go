package authorize

import (
	"users/user"

	"github.com/o1egl/paseto"
)

type (
	Checker interface {
		CheckExists(string) error
	}

	service struct {
		secret string
		r      Checker
	}
)

func NewService(secret string, r Checker) *service {
	return &service{secret, r}
}

func (s *service) Authorize(token string) (string, error) {
	if token == "" {
		return "", user.ErrMissingData
	}

	var data paseto.JSONToken

	err := paseto.NewV2().Decrypt(token, []byte(s.secret), &data, nil)
	if err != nil {
		return "", user.ErrInvalidToken
	}

	err = s.r.CheckExists(data.Jti)
	if err == nil {
		return "", user.ErrInvalidToken
	}

	userID := data.Subject

	return userID, nil
}
