package edit

import (
	"testing"

	"plants/list"
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
		findService := list.NewListService(repo)
		editService := NewEditService(findService, repo)
		id := "123"
		newPlant := plant.Plant{
			Name: "",
		}
		result, err := editService.Edit(uid, id, newPlant)

		require.EqualError(t, err, plant.ErrMissingData.Error())
		require.Equal(t, int64(0), result)
	})

	t.Run("should error when the plant is not found", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []plant.Plant{
				{
					ID:   "123",
					Name: "test",
				},
			},
		}
		findService := list.NewListService(repo)
		editService := NewEditService(findService, repo)
		id := "122"
		newPlant := plant.Plant{
			Name: "test",
		}
		result, err := editService.Edit(uid, id, newPlant)

		require.EqualError(t, err, plant.ErrNotFound.Error())
		require.Equal(t, int64(0), result)
	})

	t.Run("should edit a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []plant.Plant{
				{
					ID:   "123",
					Name: "test",
				},
			},
		}
		findService := list.NewListService(repo)
		editService := NewEditService(findService, repo)
		id := "123"
		newPlant := plant.Plant{
			Name: "test2",
		}
		result, err := editService.Edit(uid, id, newPlant)

		require.NoError(t, err)
		require.Equal(t, int64(1), result)
	})
}
