package authenticate

import (
	"testing"
	u "users/src/user"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

type mockRepo struct {
	Users []u.User
}

func (m mockRepo) InsertOne(new u.User) (interface{}, error) {
	m.Users = append(m.Users, new)

	return new, nil
}

func (m mockRepo) FindOne(email string) (u.User, error) {
	for _, u := range m.Users {
		if email == u.Email {
			return u, nil
		}
	}

	return u.User{}, u.ErrNotFound
}

func TestAuthenticate(t *testing.T) {
	t.Run("should error when there are missing fields", func(t *testing.T) {
		t.Parallel()

		repo := mockRepo{
			Users: []u.User{},
		}
		authenticateService := NewAuthenticateService(repo)
		creds := u.Credentials{
			Email: "test@test.com",
		}
		jwt, err := authenticateService.Authenticate(creds)

		require.Empty(t, jwt)
		require.EqualError(t, err, u.ErrMissingData.Error())
	})

	t.Run("should error when user does not exist", func(t *testing.T) {
		t.Parallel()

		repo := mockRepo{
			Users: []u.User{},
		}
		authenticateService := NewAuthenticateService(repo)
		creds := u.Credentials{
			Email:    "test@test.com",
			Password: "1234",
		}
		jwt, err := authenticateService.Authenticate(creds)

		require.Empty(t, jwt)
		require.EqualError(t, err, u.ErrNotFound.Error())
	})

	t.Run("should error when password is incorrect", func(t *testing.T) {
		t.Parallel()

		password := "1234"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
		repo := mockRepo{
			Users: []u.User{
				{
					Name:  "test",
					Email: "test@test.com",
					Hash:  string(hash),
				},
			},
		}
		authenticateService := NewAuthenticateService(repo)
		creds := u.Credentials{
			Email:    "test@test.com",
			Password: "12345",
		}
		jwt, err := authenticateService.Authenticate(creds)

		require.Empty(t, jwt)
		require.EqualError(t, err, u.ErrInvalidPassword.Error())
	})

	t.Run("should return a JWT on correct authentication", func(t *testing.T) {
		t.Parallel()

		password := "1234"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
		repo := mockRepo{
			Users: []u.User{
				{
					Name:  "test",
					Email: "test@test.com",
					Hash:  string(hash),
				},
			},
		}
		authenticateService := NewAuthenticateService(repo)
		creds := u.Credentials{
			Email:    "test@test.com",
			Password: password,
		}
		jwt, err := authenticateService.Authenticate(creds)

		require.NotEmpty(t, jwt)
		require.NoError(t, err)
	})
}
