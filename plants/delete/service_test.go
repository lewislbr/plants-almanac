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
		ds := NewService(repo)
		id := ""
		err := ds.Delete(uid, id)

		require.EqualError(t, err, plant.ErrMissingData.Error())
	})

	t.Run("should error when there are no matches", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []plant.Plant{
				{
					ID:   "123",
					Name: "test",
				},
			},
		}
		ds := NewService(repo)
		id := "124"
		err := ds.Delete(uid, id)

		require.EqualError(t, err, plant.ErrNotFound.Error())
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
		ds := NewService(repo)
		id := "123"
		err := ds.Delete(uid, id)

		require.NoError(t, err)
	})
}
