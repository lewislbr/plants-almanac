package generate

import (
	"testing"

	"users/user"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	t.Run("should error when the user ID is empty", func(t *testing.T) {
		t.Parallel()

		generateSvc := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V")
		userID := ""
		userID, err := generateSvc.GenerateToken(userID)

		require.Empty(t, userID)
		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should generate a token given a user ID", func(t *testing.T) {
		t.Parallel()

		generateSvc := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V")
		userID := "123"
		userID, err := generateSvc.GenerateToken(userID)

		require.NotEmpty(t, userID)
		require.NoError(t, err)
	})
}
