package revoke

import (
	"users/user"

	"github.com/o1egl/paseto"
)

type (
	Adder interface {
		Add(string) error
	}

	service struct {
		secret string
		r      Adder
	}
)

func NewService(secret string, r Adder) *service {
	return &service{secret, r}
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

	return s.r.Add(data.Jti)
}
