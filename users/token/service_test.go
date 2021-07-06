package token

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	t.Run("should error when the user ID is empty", func(t *testing.T) {
		t.Parallel()

		redis := &mockRedisRepository{}
		service := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", redis)
		userID := ""
		userID, err := service.Generate(userID)

		require.Empty(t, userID)
		require.ErrorIs(t, err, ErrMissingData)
	})

	t.Run("should generate a token given a user ID", func(t *testing.T) {
		t.Parallel()

		redis := &mockRedisRepository{}
		service := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", redis)
		userID := "123"
		userID, err := service.Generate(userID)

		require.NotEmpty(t, userID)
		require.NoError(t, err)
	})
}

func TestValidate(t *testing.T) {
	t.Run("should error when token is empty", func(t *testing.T) {
		t.Parallel()

		redis := &mockRedisRepository{}

		redis.On("CheckExists", mock.AnythingOfType("string")).Return(errors.New("not found"))

		service := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", redis)
		token := ""
		userID, err := service.Validate(token)

		require.Empty(t, userID)
		require.ErrorIs(t, err, ErrMissingData)
	})

	t.Run("should error when token is invalid", func(t *testing.T) {
		t.Parallel()

		redis := &mockRedisRepository{}

		redis.On("CheckExists", mock.AnythingOfType("string")).Return(errors.New("not found"))

		service := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", redis)
		token := "a.b.c.d"
		userID, err := service.Validate(token)

		require.Empty(t, userID)
		require.ErrorIs(t, err, ErrInvalidToken)
	})

	t.Run("should error when token is revoked", func(t *testing.T) {
		t.Parallel()

		redis := &mockRedisRepository{}

		redis.On("CheckExists", mock.AnythingOfType("string")).Return(nil)

		service := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", redis)
		token := "a.b.c.d"
		userID, err := service.Validate(token)

		require.Empty(t, userID)
		require.ErrorIs(t, err, ErrInvalidToken)
	})

	t.Run("should return an ID", func(t *testing.T) {
		t.Parallel()

		redis := &mockRedisRepository{}

		redis.On("CheckExists", mock.AnythingOfType("string")).Return(errors.New("not found"))

		service := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", redis)
		expectedID := "123"
		token := "v2.local.y4IJ_w7Sn6FTFdRbtzhVkSHg85QX7kSUiyKofqHtoSm-6rGh9HwJikea1mhuYAAAzbk0UHa5O5SGLl2Ztc6udGtcuuxo9diBC0VqgZ34sRuaZWgy0JypVOqntXvvApo7QcE4AUjO3wimRtzJMbgexLXKvV6xgWwrnDGQvYK2pKBG1ww-7YNmCSkEK6YuxOF3eefvrVr5D3E4gJNNAXvQSx1vrVlr82GlTmy2z29F-QrmD1-m6phxYAiKTQ.bnVsbA" // Token with no expiration
		userID, err := service.Validate(token)

		require.Equal(t, expectedID, userID)
		require.NoError(t, err)
	})
}

func TestRevoke(t *testing.T) {
	t.Run("should error when the token is empty", func(t *testing.T) {
		t.Parallel()

		redis := &mockRedisRepository{}

		redis.On("Add", mock.AnythingOfType("string")).Return(nil)

		service := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", redis)
		token := ""
		err := service.Revoke(token)

		require.ErrorIs(t, err, ErrMissingData)
	})

	t.Run("should error when the token is invalid", func(t *testing.T) {
		t.Parallel()

		redis := &mockRedisRepository{}

		redis.On("Add", mock.AnythingOfType("string")).Return(nil)

		service := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", redis)
		token := "a.b.c.d"
		err := service.Revoke(token)

		require.ErrorIs(t, err, ErrInvalidToken)
	})

	t.Run("should return no error on success", func(t *testing.T) {
		t.Parallel()

		redis := &mockRedisRepository{}

		redis.On("Add", mock.AnythingOfType("string")).Return(nil)

		service := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", redis)
		token := "v2.local.y4IJ_w7Sn6FTFdRbtzhVkSHg85QX7kSUiyKofqHtoSm-6rGh9HwJikea1mhuYAAAzbk0UHa5O5SGLl2Ztc6udGtcuuxo9diBC0VqgZ34sRuaZWgy0JypVOqntXvvApo7QcE4AUjO3wimRtzJMbgexLXKvV6xgWwrnDGQvYK2pKBG1ww-7YNmCSkEK6YuxOF3eefvrVr5D3E4gJNNAXvQSx1vrVlr82GlTmy2z29F-QrmD1-m6phxYAiKTQ.bnVsbA"
		err := service.Revoke(token)

		require.NoError(t, err)
	})
}
