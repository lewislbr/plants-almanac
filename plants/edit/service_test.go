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
		err := editService.Edit(uid, id, newPlant)

		require.EqualError(t, err, plant.ErrMissingData.Error())
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
		err := editService.Edit(uid, id, newPlant)

		require.EqualError(t, err, plant.ErrNotFound.Error())
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
		err := editService.Edit(uid, id, newPlant)

		require.NoError(t, err)
	})
}
