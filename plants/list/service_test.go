package list

import (
	"testing"

	p "plants/plant"
	"plants/storage"

	"github.com/stretchr/testify/require"
)

const uid = "123"

func TestCreate(t *testing.T) {
	t.Run("should list all plants with no error", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []p.Plant{
				{
					ID:   "123",
					Name: "test",
				},
			},
		}
		listService := NewListService(repo)
		result, err := listService.ListAll(uid)

		require.NoError(t, err)
		require.NotNil(t, result)
	})

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
		listService := NewListService(repo)
		id := ""
		result, err := listService.ListOne(uid, id)

		require.EqualError(t, err, p.ErrMissingData.Error())
		require.Equal(t, p.Plant{}, result)
	})

	t.Run("should error when a plant is not found", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []p.Plant{
				{
					ID:   "123",
					Name: "test",
				},
			},
		}
		listService := NewListService(repo)
		id := "122"
		result, err := listService.ListOne(uid, id)

		require.EqualError(t, err, p.ErrNotFound.Error())
		require.Equal(t, p.Plant{}, result)
	})

	t.Run("should list a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []p.Plant{
				{
					ID:   "123",
					Name: "test",
				},
			},
		}
		listService := NewListService(repo)
		id := "123"
		result, err := listService.ListOne(uid, id)

		require.NoError(t, err)
		require.NotNil(t, result)
	})
}
