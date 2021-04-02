package create

import (
	"testing"

	"users/storage"
	"users/user"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Run("should error when there are missing fields", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Users: []user.User{},
		}
		cs := NewService(repo)
		new := user.User{
			Email:    "test@test.com",
			Password: "1234",
		}
		err := cs.Create(new)

		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should error when the user already exists", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Users: []user.User{
				{
					Name:     "test",
					Email:    "test@test.com",
					Password: "1234",
				},
			},
		}
		cs := NewService(repo)
		new := user.User{
			Name:     "test",
			Email:    "test@test.com",
			Password: "1234",
		}
		err := cs.Create(new)

		require.EqualError(t, err, user.ErrUserExists.Error())
	})

	t.Run("should create a user with no error", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Users: []user.User{},
		}
		cs := NewService(repo)
		new := user.User{
			Name:     "test",
			Email:    "test@test.com",
			Password: "1234",
		}
		err := cs.Create(new)

		require.NoError(t, err)
	})
}
