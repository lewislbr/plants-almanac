package authorize

import (
	"errors"
	"testing"

	"users/user"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAuthorization(t *testing.T) {
	t.Run("should error when token is empty", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("CheckExists", mock.AnythingOfType("string")).Return(errors.New("not found"))

		authorizeSvc := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", repo)
		token := ""
		userID, err := authorizeSvc.Authorize(token)

		require.Empty(t, userID)
		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should error when token is invalid", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("CheckExists", mock.AnythingOfType("string")).Return(errors.New("not found"))

		authorizeSvc := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", repo)
		token := "a.b.c.d"
		userID, err := authorizeSvc.Authorize(token)

		require.Empty(t, userID)
		require.EqualError(t, err, user.ErrInvalidToken.Error())
	})

	t.Run("should error when token is revoked", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("CheckExists", mock.AnythingOfType("string")).Return(nil)

		authorizeSvc := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", repo)
		token := "a.b.c.d"
		userID, err := authorizeSvc.Authorize(token)

		require.Empty(t, userID)
		require.EqualError(t, err, user.ErrInvalidToken.Error())
	})

	t.Run("should return an ID", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("CheckExists", mock.AnythingOfType("string")).Return(errors.New("not found"))

		authorizeSvc := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", repo)
		expectedID := "123"
		token := "v2.local.y4IJ_w7Sn6FTFdRbtzhVkSHg85QX7kSUiyKofqHtoSm-6rGh9HwJikea1mhuYAAAzbk0UHa5O5SGLl2Ztc6udGtcuuxo9diBC0VqgZ34sRuaZWgy0JypVOqntXvvApo7QcE4AUjO3wimRtzJMbgexLXKvV6xgWwrnDGQvYK2pKBG1ww-7YNmCSkEK6YuxOF3eefvrVr5D3E4gJNNAXvQSx1vrVlr82GlTmy2z29F-QrmD1-m6phxYAiKTQ.bnVsbA" // Token with no expiration
		userID, err := authorizeSvc.Authorize(token)

		require.Equal(t, expectedID, userID)
		require.NoError(t, err)
	})
}
