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
		ls := list.NewService(repo)
		es := NewService(ls, repo)
		id := "123"
		newPlant := plant.Plant{
			Name: "",
		}
		err := es.Edit(uid, id, newPlant)

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
		ls := list.NewService(repo)
		es := NewService(ls, repo)
		id := "122"
		newPlant := plant.Plant{
			Name: "test",
		}
		err := es.Edit(uid, id, newPlant)

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
		ls := list.NewService(repo)
		es := NewService(ls, repo)
		id := "123"
		newPlant := plant.Plant{
			Name: "test2",
		}
		err := es.Edit(uid, id, newPlant)

		require.NoError(t, err)
	})
}
