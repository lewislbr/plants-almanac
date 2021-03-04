package edit

import (
	"testing"

	"plants/list"
	p "plants/plant"
	"plants/storage"

	"github.com/stretchr/testify/require"
)

const uid = "123"

func TestCreate(t *testing.T) {
	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []p.Plant{
				{
					ID:   "123",
					Name: "test",
				},
			},
		}
		findService := list.NewListService(repo)
		editService := NewEditService(findService, repo)
		id := "123"
		newPlant := p.Plant{
			Name: "",
		}
		result, err := editService.Edit(uid, id, newPlant)

		require.EqualError(t, err, p.ErrMissingData.Error())
		require.Equal(t, int64(0), result)
	})

	t.Run("should error when the plant is not found", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []p.Plant{
				{
					ID:   "123",
					Name: "test",
				},
			},
		}
		findService := list.NewListService(repo)
		editService := NewEditService(findService, repo)
		id := "122"
		newPlant := p.Plant{
			Name: "test",
		}
		result, err := editService.Edit(uid, id, newPlant)

		require.EqualError(t, err, p.ErrNotFound.Error())
		require.Equal(t, int64(0), result)
	})

	t.Run("should edit a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []p.Plant{
				{
					ID:   "123",
					Name: "test",
				},
			},
		}
		findService := list.NewListService(repo)
		editService := NewEditService(findService, repo)
		id := "123"
		newPlant := p.Plant{
			Name: "test2",
		}
		result, err := editService.Edit(uid, id, newPlant)

		require.NoError(t, err)
		require.Equal(t, int64(1), result)
	})
}
