package token

import (
	"fmt"
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
		return "", fmt.Errorf("error generating token: %w", ErrMissingData)
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
		return "", fmt.Errorf("error encrypting token: %w", err)
	}

	return token, nil
}

func (s *service) Validate(token string) (string, error) {
	if token == "" {
		return "", fmt.Errorf("error validating token: %w", ErrMissingData)
	}

	var data paseto.JSONToken

	err := paseto.NewV2().Decrypt(token, []byte(s.secret), &data, nil)
	if err != nil {
		return "", fmt.Errorf("error decrypting token: %w", ErrInvalidToken)
	}

	err = s.redis.CheckExists(data.Jti)
	if err == nil {
		return "", fmt.Errorf("error checking token: %w", ErrInvalidToken)
	}

	userID := data.Subject

	return userID, nil
}

func (s *service) Revoke(token string) error {
	if token == "" {
		return fmt.Errorf("error revoking token: %w", ErrMissingData)
	}

	var data paseto.JSONToken

	err := paseto.NewV2().Decrypt(token, []byte(s.secret), &data, nil)
	if err != nil {
		return fmt.Errorf("error decrypting token: %w", ErrInvalidToken)
	}

	err = s.redis.Add(data.Jti)
	if err != nil {
		return fmt.Errorf("error adding token: %w", err)
	}

	return nil
}
