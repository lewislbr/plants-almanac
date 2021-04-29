package delete

import (
	"testing"

	"plants/plant"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const userID = "123"

func TestCreate(t *testing.T) {
	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("DeleteOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(1), nil)

		deleteSvc := NewService(repo)
		plantID := ""
		err := deleteSvc.Delete(userID, plantID)

		require.EqualError(t, err, plant.ErrMissingData.Error())
	})

	t.Run("should error when there are no matches", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("DeleteOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(0), nil)

		deleteSvc := NewService(repo)
		plantID := "124"
		err := deleteSvc.Delete(userID, plantID)

		require.EqualError(t, err, plant.ErrNotFound.Error())
	})

	t.Run("should delete a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("DeleteOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(1), nil)

		deleteSvc := NewService(repo)
		plantID := "123"
		err := deleteSvc.Delete(userID, plantID)

		require.NoError(t, err)
	})
}
