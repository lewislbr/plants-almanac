package authorize

import (
	"testing"

	"users/user"

	"github.com/stretchr/testify/require"
)

func TestAuthorization(t *testing.T) {
	t.Run("should error when JWT is empty", func(t *testing.T) {
		t.Parallel()

		authorizeService := NewAuthorizeService("test")
		jwt := ""
		id, err := authorizeService.Authorize(jwt)

		require.Empty(t, id)
		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should error when JWT is invalid", func(t *testing.T) {
		t.Parallel()

		authorizeService := NewAuthorizeService("test")
		jwt := "a.b.c"
		id, err := authorizeService.Authorize(jwt)

		require.Empty(t, id)
		require.EqualError(t, err, user.ErrInvalidToken.Error())
	})

	t.Run("should return an ID", func(t *testing.T) {
		t.Parallel()

		authorizeService := NewAuthorizeService("test")
		expectedID := "1"
		jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIn0.bHGULc9qoVyle089kDZjXUBqDzsFHjRO074sqv_ILW8" // Token with no expiration
		id, err := authorizeService.Authorize(jwt)

		require.Equal(t, expectedID, id)
		require.NoError(t, err)
	})
}
