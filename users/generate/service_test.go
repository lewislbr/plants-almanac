package generate

import (
	"testing"

	"users/user"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	t.Run("should error when uid is empty", func(t *testing.T) {
		t.Parallel()

		gs := NewService("test")
		uid := ""
		id, err := gs.GenerateToken(uid)

		require.Empty(t, id)
		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should generate a token given a uid", func(t *testing.T) {
		t.Parallel()

		gs := NewService("WNxmZvttwv2YmvS3JWqpJ6vNd3YpQw6V")
		uid := "123"
		id, err := gs.GenerateToken(uid)

		require.NotEmpty(t, id)
		require.NoError(t, err)
	})
}
