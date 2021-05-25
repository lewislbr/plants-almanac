package token

import (
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type (
	redisRepository interface {
		Add(string) error
		CheckExists(string) error
	}

	service struct {
		secret string
		redis  redisRepository
	}
)

func NewService(secret string, redis redisRepository) *service {
	return &service{secret, redis}
}

func (s *service) Generate(userID string) (string, error) {
	if userID == "" {
		return "", ErrMissingData
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

func (s *service) Validate(token string) (string, error) {
	if token == "" {
		return "", ErrMissingData
	}

	var data paseto.JSONToken

	err := paseto.NewV2().Decrypt(token, []byte(s.secret), &data, nil)
	if err != nil {
		return "", ErrInvalidToken
	}

	err = s.redis.CheckExists(data.Jti)
	if err == nil {
		return "", ErrInvalidToken
	}

	userID := data.Subject

	return userID, nil
}

func (s *service) Revoke(token string) error {
	if token == "" {
		return ErrMissingData
	}

	var data paseto.JSONToken

	err := paseto.NewV2().Decrypt(token, []byte(s.secret), &data, nil)
	if err != nil {
		return ErrInvalidToken
	}

	return s.redis.Add(data.Jti)
}
