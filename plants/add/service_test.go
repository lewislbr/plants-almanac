package add

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
			Plants: []plant.Plant{},
		}
		as := NewService(repo)
		newPlant := plant.Plant{
			Name: "",
		}
		err := as.Add(uid, newPlant)

		require.EqualError(t, err, plant.ErrMissingData.Error())
	})

	t.Run("should create a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []plant.Plant{},
		}
		as := NewService(repo)
		newPlant := plant.Plant{
			Name: "test",
		}
		err := as.Add(uid, newPlant)

		require.NoError(t, err)
	})
}
