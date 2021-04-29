package authenticate

import (
	"errors"
	"testing"

	"users/user"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticate(t *testing.T) {
	t.Run("should error when there are missing fields", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindOne", mock.AnythingOfType("string")).Return(user.User{}, nil)

		authenticateSvc := NewService(repo)
		creds := user.Credentials{
			Email: "test@test.com",
		}
		userID, err := authenticateSvc.Authenticate(creds)

		require.Empty(t, userID)
		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should error when the user does not exist", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindOne", mock.AnythingOfType("string")).Return(user.User{}, errors.New("user not found"))

		authenticateSvc := NewService(repo)
		creds := user.Credentials{
			Email:    "test@test.com",
			Password: "123",
		}
		userID, err := authenticateSvc.Authenticate(creds)

		require.Empty(t, userID)
		require.EqualError(t, err, user.ErrNotFound.Error())
	})

	t.Run("should error when password is incorrect", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}
		password := "123"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

		repo.On("FindOne", mock.AnythingOfType("string")).Return(
			user.User{
				ID:    "1",
				Name:  "test",
				Email: "test@test.com",
				Hash:  string(hash),
			},
			nil,
		)

		authenticateSvc := NewService(repo)
		creds := user.Credentials{
			Email:    "test@test.com",
			Password: "321",
		}
		userID, err := authenticateSvc.Authenticate(creds)

		require.Empty(t, userID)
		require.EqualError(t, err, user.ErrInvalidPassword.Error())
	})

	t.Run("should return a user userID on correct authentication", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}
		password := "123"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

		repo.On("FindOne", mock.AnythingOfType("string")).Return(
			user.User{
				ID:    "123",
				Name:  "test",
				Email: "test@test.com",
				Hash:  string(hash),
			},
			nil,
		)

		authenticateSvc := NewService(repo)
		creds := user.Credentials{
			Email:    "test@test.com",
			Password: password,
		}
		userID, err := authenticateSvc.Authenticate(creds)

		require.NotEmpty(t, userID)
		require.NoError(t, err)
	})
}
