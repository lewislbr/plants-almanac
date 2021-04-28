package info

import (
	"testing"

	"users/user"

	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestInfo(t *testing.T) {
	t.Run("should error when id is empty", func(t *testing.T) {
		t.Parallel()

		g := &MockGetter{}

		g.On("GetUserInfo", mock.AnythingOfType("string")).Return(user.Info{}, nil)

		is := NewService(g)
		id := ""
		_, err := is.UserInfo(id)

		require.Empty(t, id)
		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should return user info when request is successful", func(t *testing.T) {
		t.Parallel()

		g := &MockGetter{}

		g.On("GetUserInfo", mock.AnythingOfType("string")).Return(user.Info{}, nil)

		is := NewService(g)
		id := "123"
		_, err := is.UserInfo(id)

		require.NoError(t, err)
	})
}
