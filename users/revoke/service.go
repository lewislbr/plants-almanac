package revoke

import (
	"users/user"

	"github.com/o1egl/paseto"
)

type (
	repository interface {
		Add(string) error
	}

	service struct {
		secret string
		repo   repository
	}
)

func NewService(secret string, repo repository) *service {
	return &service{secret, repo}
}

func (s *service) RevokeToken(token string) error {
	if token == "" {
		return user.ErrMissingData
	}

	var data paseto.JSONToken

	err := paseto.NewV2().Decrypt(token, []byte(s.secret), &data, nil)
	if err != nil {
		return user.ErrInvalidToken
	}

	return s.repo.Add(data.Jti)
}
