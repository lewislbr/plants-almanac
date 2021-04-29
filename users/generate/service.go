package generate

import (
	"time"

	"users/user"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type service struct {
	secret string
}

func NewService(secret string) *service {
	return &service{secret}
}

func (s *service) GenerateToken(userID string) (string, error) {
	if userID == "" {
		return "", user.ErrMissingData
	}

	jsonToken := paseto.JSONToken{
		Audience:   "plantdex",
		Issuer:     "users",
		Jti:        uuid.New().String(),
		Subject:    userID,
		IssuedAt:   time.Now(),
		Expiration: time.Now().AddDate(0, 0, 7),
		NotBefore:  time.Now(),
	}
	token, err := paseto.NewV2().Encrypt([]byte(s.secret), jsonToken, nil)
	if err != nil {
		return "", err
	}

	return token, nil
}
