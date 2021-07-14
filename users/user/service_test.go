package user

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestCreate(t *testing.T) {
	t.Run("should error when there are missing fields", func(t *testing.T) {
		t.Parallel()

		userRepo := &mockUserRepo{}

		userRepo.On("CheckExists", mock.AnythingOfType("string")).Return(false, nil)
		userRepo.On("Insert", mock.AnythingOfType("User")).Return(nil)

		userService := NewService(userRepo)
		new := New{
			Email:    "test@test.com",
			Password: "123",
		}
		err := userService.Create(new)

		require.ErrorIs(t, err, ErrMissingData)
	})

	t.Run("should error when the user already exists", func(t *testing.T) {
		t.Parallel()

		userRepo := &mockUserRepo{}

		userRepo.On("CheckExists", mock.AnythingOfType("string")).Return(true, nil)
		userRepo.On("Insert", mock.AnythingOfType("User")).Return(nil)

		userService := NewService(userRepo)
		new := New{
			Name:     "test",
			Email:    "test@test.com",
			Password: "123",
		}
		err := userService.Create(new)

		require.ErrorIs(t, err, ErrUserExists)
	})

	t.Run("should create a user with no error", func(t *testing.T) {
		t.Parallel()

		userRepo := &mockUserRepo{}

		userRepo.On("CheckExists", mock.AnythingOfType("string")).Return(false, nil)
		userRepo.On("Insert", mock.AnythingOfType("User")).Return(nil)

		userService := NewService(userRepo)
		new := New{
			Name:     "test",
			Email:    "test@test.com",
			Password: "123",
		}
		err := userService.Create(new)

		require.NoError(t, err)
	})
}

func TestAuthenticate(t *testing.T) {
	t.Run("should error when there are missing fields", func(t *testing.T) {
		t.Parallel()

		userRepo := &mockUserRepo{}

		userRepo.On("Find", mock.AnythingOfType("string")).Return(User{}, nil)

		userService := NewService(userRepo)
		creds := Credentials{
			Email: "test@test.com",
		}
		userID, err := userService.Authenticate(creds)

		require.Empty(t, userID)
		require.ErrorIs(t, err, ErrMissingData)
	})

	t.Run("should error when the user does not exist", func(t *testing.T) {
		t.Parallel()

		userRepo := &mockUserRepo{}

		userRepo.On("Find", mock.AnythingOfType("string")).Return(User{}, nil)

		userService := NewService(userRepo)
		creds := Credentials{
			Email:    "test@test.com",
			Password: "123",
		}
		userID, err := userService.Authenticate(creds)

		require.Empty(t, userID)
		require.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("should error when password is incorrect", func(t *testing.T) {
		t.Parallel()

		userRepo := &mockUserRepo{}
		password := "123"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

		userRepo.On("Find", mock.AnythingOfType("string")).Return(
			User{
				ID:    "1",
				Name:  "test",
				Email: "test@test.com",
				Hash:  string(hash),
			},
			nil,
		)

		userService := NewService(userRepo)
		creds := Credentials{
			Email:    "test@test.com",
			Password: "321",
		}
		userID, err := userService.Authenticate(creds)

		require.Empty(t, userID)
		require.ErrorIs(t, err, ErrInvalidPassword)
	})

	t.Run("should return a user userID on correct authentication", func(t *testing.T) {
		t.Parallel()

		userRepo := &mockUserRepo{}
		password := "123"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

		userRepo.On("Find", mock.AnythingOfType("string")).Return(
			User{
				ID:    "123",
				Name:  "test",
				Email: "test@test.com",
				Hash:  string(hash),
			},
			nil,
		)

		userService := NewService(userRepo)
		creds := Credentials{
			Email:    "test@test.com",
			Password: password,
		}
		userID, err := userService.Authenticate(creds)

		require.NotEmpty(t, userID)
		require.NoError(t, err)
	})
}

func TestInfo(t *testing.T) {
	t.Run("should error when the user ID is empty", func(t *testing.T) {
		t.Parallel()

		userRepo := &mockUserRepo{}

		userRepo.On("GetInfo", mock.AnythingOfType("string")).Return(Info{}, nil)

		userService := NewService(userRepo)
		userID := ""
		_, err := userService.Info(userID)

		require.Empty(t, userID)
		require.ErrorIs(t, err, ErrMissingData)
	})

	t.Run("should return user info when request is successful", func(t *testing.T) {
		t.Parallel()

		userRepo := &mockUserRepo{}

		userRepo.On("GetInfo", mock.AnythingOfType("string")).Return(Info{Name: "foo"}, nil)

		userService := NewService(userRepo)
		userID := "123"
		_, err := userService.Info(userID)

		require.NoError(t, err)
	})
}
