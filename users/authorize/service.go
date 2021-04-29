package authorize

import (
	"users/user"

	"github.com/o1egl/paseto"
)

type (
	repository interface {
		CheckExists(string) error
	}

	service struct {
		secret string
		repo   repository
	}
)

func NewService(secret string, repo repository) *service {
	return &service{secret, repo}
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

	err = s.repo.CheckExists(data.Jti)
	if err == nil {
		return "", user.ErrInvalidToken
	}

	userID := data.Subject

	return userID, nil
}
