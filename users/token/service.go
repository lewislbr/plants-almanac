package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

const expiration = 30 * 24 * time.Hour

type (
	tokenRepo interface {
		Add(string) error
		CheckExists(string) (bool, error)
	}

	tokenService struct {
		secret    string
		tokenRepo tokenRepo
	}
)

func NewService(secret string, tokenRepo tokenRepo) *tokenService {
	return &tokenService{secret, tokenRepo}
}

func (s *tokenService) Generate(userID string) (string, error) {
	if userID == "" {
		return "", fmt.Errorf("error generating token: %w", ErrMissingData)
	}

	now := time.Now().UTC()
	jsonToken := paseto.JSONToken{
		Audience:   "plantdex",
		Issuer:     "users",
		Jti:        uuid.New().String(),
		Subject:    userID,
		IssuedAt:   now,
		Expiration: now.Add(expiration),
		NotBefore:  now,
	}
	token, err := paseto.NewV2().Encrypt([]byte(s.secret), jsonToken, nil)
	if err != nil {
		return "", fmt.Errorf("error encrypting token: %w", err)
	}

	return token, nil
}

func (s *tokenService) Validate(token string) (string, error) {
	if token == "" {
		return "", fmt.Errorf("error validating token: %w", ErrMissingData)
	}

	var data paseto.JSONToken

	err := paseto.NewV2().Decrypt(token, []byte(s.secret), &data, nil)
	if err != nil {
		return "", fmt.Errorf("error decrypting token: %w", ErrInvalidToken)
	}

	exists, err := s.tokenRepo.CheckExists(data.Jti)
	if err != nil {
		return "", fmt.Errorf("error checking token: %w", err)
	}
	if exists {
		return "", fmt.Errorf("error checking token: %w", ErrInvalidToken)
	}

	userID := data.Subject

	return userID, nil
}

func (s *tokenService) Revoke(token string) error {
	if token == "" {
		return fmt.Errorf("error revoking token: %w", ErrMissingData)
	}

	var data paseto.JSONToken

	err := paseto.NewV2().Decrypt(token, []byte(s.secret), &data, nil)
	if err != nil {
		return fmt.Errorf("error decrypting token: %w", ErrInvalidToken)
	}

	err = s.tokenRepo.Add(data.Jti)
	if err != nil {
		return fmt.Errorf("error adding token: %w", err)
	}

	return nil
}
