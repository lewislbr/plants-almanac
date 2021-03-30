package delete

import (
	"testing"

	"plants/plant"
	"plants/storage"

	"github.com/stretchr/testify/require"
)

const uid = "123"

func TestCreate(t *testing.T) {
	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []plant.Plant{
				{
					ID:   "123",
					Name: "test",
				},
			},
		}
		deleteService := NewDeleteService(repo)
		id := ""
		result, err := deleteService.Delete(uid, id)

		require.EqualError(t, err, plant.ErrMissingData.Error())
		require.Equal(t, int64(0), result)
	})

	t.Run("should delete a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []plant.Plant{
				{
					ID:   "123",
					Name: "test",
				},
			},
		}
		deleteService := NewDeleteService(repo)
		id := "123"
		result, err := deleteService.Delete(uid, id)

		require.NoError(t, err)
		require.Equal(t, int64(1), result)
	})
}
