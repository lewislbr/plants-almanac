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

		i := &MockInserterChecker{}

		i.On("CheckExists", mock.AnythingOfType("string")).Return(false, nil)
		i.On("InsertOne", mock.AnythingOfType("user.User")).Return(nil)

		cs := NewService(i)
		new := user.User{
			Email:    "test@test.com",
			Password: "123",
		}
		err := cs.Create(new)

		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should error when the user already exists", func(t *testing.T) {
		t.Parallel()

		i := &MockInserterChecker{}

		i.On("CheckExists", mock.AnythingOfType("string")).Return(true, nil)
		i.On("InsertOne", mock.AnythingOfType("user.User")).Return(nil)

		cs := NewService(i)
		new := user.User{
			Name:     "test",
			Email:    "test@test.com",
			Password: "123",
		}
		err := cs.Create(new)

		require.EqualError(t, err, user.ErrUserExists.Error())
	})

	t.Run("should create a user with no error", func(t *testing.T) {
		t.Parallel()

		i := &MockInserterChecker{}

		i.On("CheckExists", mock.AnythingOfType("string")).Return(false, nil)
		i.On("InsertOne", mock.AnythingOfType("user.User")).Return(nil)

		cs := NewService(i)
		new := user.User{
			Name:     "test",
			Email:    "test@test.com",
			Password: "123",
		}
		err := cs.Create(new)

		require.NoError(t, err)
	})
}
