package info

import (
	"testing"

	"users/user"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestInfo(t *testing.T) {
	t.Run("should error when the user ID is empty", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("GetUserInfo", mock.AnythingOfType("string")).Return(user.Info{}, nil)

		infoSvc := NewService(repo)
		userID := ""
		_, err := infoSvc.UserInfo(userID)

		require.Empty(t, userID)
		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should return user info when request is successful", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("GetUserInfo", mock.AnythingOfType("string")).Return(user.Info{}, nil)

		infoSvc := NewService(repo)
		userID := "123"
		_, err := infoSvc.UserInfo(userID)

		require.NoError(t, err)
	})
}
