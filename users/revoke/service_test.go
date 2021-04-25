package revoke

import (
	"testing"
	"users/user"

	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	t.Run("should error when the token is empty", func(t *testing.T) {
		t.Parallel()

		a := &MockAdder{}

		a.On("Add", mock.AnythingOfType("string")).Return(nil)

		rs := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", a)
		token := ""
		err := rs.RevokeToken(token)

		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should error when the token is invalid", func(t *testing.T) {
		t.Parallel()

		a := &MockAdder{}

		a.On("Add", mock.AnythingOfType("string")).Return(nil)

		rs := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", a)
		token := "a.b.c.d"
		err := rs.RevokeToken(token)

		require.EqualError(t, err, user.ErrInvalidToken.Error())
	})

	t.Run("should return no error on success", func(t *testing.T) {
		t.Parallel()

		a := &MockAdder{}

		a.On("Add", mock.AnythingOfType("string")).Return(nil)

		rs := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V", a)
		token := "v2.local.y4IJ_w7Sn6FTFdRbtzhVkSHg85QX7kSUiyKofqHtoSm-6rGh9HwJikea1mhuYAAAzbk0UHa5O5SGLl2Ztc6udGtcuuxo9diBC0VqgZ34sRuaZWgy0JypVOqntXvvApo7QcE4AUjO3wimRtzJMbgexLXKvV6xgWwrnDGQvYK2pKBG1ww-7YNmCSkEK6YuxOF3eefvrVr5D3E4gJNNAXvQSx1vrVlr82GlTmy2z29F-QrmD1-m6phxYAiKTQ.bnVsbA"
		err := rs.RevokeToken(token)

		require.NoError(t, err)
	})
}
