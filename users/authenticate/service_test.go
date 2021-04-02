package authenticate

import (
	"testing"

	"users/generate"
	"users/storage"
	"users/user"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticate(t *testing.T) {
	t.Run("should error when there are missing fields", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Users: []user.User{},
		}
		generateService := generate.NewGenerateService("test")
		authenticateService := NewAuthenticateService(generateService, repo)
		creds := user.Credentials{
			Email: "test@test.com",
		}
		jwt, err := authenticateService.Authenticate(creds)

		require.Empty(t, jwt)
		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should error when user does not exist", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Users: []user.User{},
		}
		generateService := generate.NewGenerateService("test")
		authenticateService := NewAuthenticateService(generateService, repo)
		creds := user.Credentials{
			Email:    "test@test.com",
			Password: "1234",
		}
		jwt, err := authenticateService.Authenticate(creds)

		require.Empty(t, jwt)
		require.EqualError(t, err, user.ErrNotFound.Error())
	})

	t.Run("should error when password is incorrect", func(t *testing.T) {
		t.Parallel()

		password := "1234"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
		repo := &storage.MockRepo{
			Users: []user.User{
				{
					Name:  "test",
					Email: "test@test.com",
					Hash:  string(hash),
				},
			},
		}
		generateService := generate.NewGenerateService("test")
		authenticateService := NewAuthenticateService(generateService, repo)
		creds := user.Credentials{
			Email:    "test@test.com",
			Password: "12345",
		}
		jwt, err := authenticateService.Authenticate(creds)

		require.Empty(t, jwt)
		require.EqualError(t, err, user.ErrInvalidPassword.Error())
	})

	t.Run("should return a JWT on correct authentication", func(t *testing.T) {
		t.Parallel()

		password := "1234"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
		repo := &storage.MockRepo{
			Users: []user.User{
				{
					ID:    "1",
					Name:  "test",
					Email: "test@test.com",
					Hash:  string(hash),
				},
			},
		}
		generateService := generate.NewGenerateService("test")
		authenticateService := NewAuthenticateService(generateService, repo)
		creds := user.Credentials{
			Email:    "test@test.com",
			Password: password,
		}
		jwt, err := authenticateService.Authenticate(creds)

		require.NotEmpty(t, jwt)
		require.NoError(t, err)
	})
}
