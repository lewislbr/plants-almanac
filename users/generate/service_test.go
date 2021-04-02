package generate

import (
	"testing"

	"users/user"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	t.Run("should error when uid is empty", func(t *testing.T) {
		t.Parallel()

		generateService := NewGenerateService("test")
		uid := ""
		id, err := generateService.GenerateJWT(uid)

		require.Empty(t, id)
		require.EqualError(t, err, user.ErrMissingData.Error())
	})

	t.Run("should generate a JWT given a uid", func(t *testing.T) {
		t.Parallel()

		generateService := NewGenerateService("test")
		uid := "1"
		id, err := generateService.GenerateJWT(uid)

		require.NotEmpty(t, id)
		require.NoError(t, err)
	})
}
