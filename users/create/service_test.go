package create

import (
	"testing"

	"users/user"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Run("should error when there are missing fields", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("CheckExists", mock.AnythingOfType("string")).Return(false, nil)
		repo.On("InsertOne", mock.AnythingOfType("user.User")).Return(nil)

		createSvc := NewService(repo)
		new := user.User{
			Email:    "test@test.com",
			Password: "123",
		}
		err := createSvc.Create(new)

		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should error when the user already exists", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("CheckExists", mock.AnythingOfType("string")).Return(true, nil)
		repo.On("InsertOne", mock.AnythingOfType("user.User")).Return(nil)

		createSvc := NewService(repo)
		new := user.User{
			Name:     "test",
			Email:    "test@test.com",
			Password: "123",
		}
		err := createSvc.Create(new)

		require.EqualError(t, err, user.ErrUserExists.Error())
	})

	t.Run("should create a user with no error", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("CheckExists", mock.AnythingOfType("string")).Return(false, nil)
		repo.On("InsertOne", mock.AnythingOfType("user.User")).Return(nil)

		createSvc := NewService(repo)
		new := user.User{
			Name:     "test",
			Email:    "test@test.com",
			Password: "123",
		}
		err := createSvc.Create(new)

		require.NoError(t, err)
	})
}
