package authorize

import (
	"users/user"

	"github.com/o1egl/paseto"
)

type service struct {
	secret string
}

func NewService(secret string) *service {
	return &service{
		secret: secret,
	}
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

	userID := data.Subject

	return userID, nil
}
