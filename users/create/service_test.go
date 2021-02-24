package create

import (
	"testing"

	"users/storage"
	u "users/user"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Run("should error when there are missing fields", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Users: []u.User{},
		}
		createService := NewCreateService(repo)
		newUser := u.User{
			Email:    "test@test.com",
			Password: "1234",
		}
		err := createService.Create(newUser)

		require.EqualError(t, err, u.ErrMissingData.Error())
	})

	t.Run("should error when the user already exists", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Users: []u.User{
				{
					Name:     "test",
					Email:    "test@test.com",
					Password: "1234",
				},
			},
		}
		createService := NewCreateService(repo)
		newUser := u.User{
			Name:     "test",
			Email:    "test@test.com",
			Password: "1234",
		}
		err := createService.Create(newUser)

		require.EqualError(t, err, u.ErrUserExists.Error())
	})

	t.Run("should create a user with no error", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Users: []u.User{},
		}
		createService := NewCreateService(repo)
		newUser := u.User{
			Name:     "test",
			Email:    "test@test.com",
			Password: "1234",
		}
		err := createService.Create(newUser)

		require.NoError(t, err)
	})
}
