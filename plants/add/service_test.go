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
		addService := NewAddService(repo)
		newPlant := plant.Plant{
			Name: "",
		}
		result, err := addService.Add(uid, newPlant)

		require.EqualError(t, err, plant.ErrMissingData.Error())
		require.Nil(t, result)
	})

	t.Run("should create a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &storage.MockRepo{
			Plants: []plant.Plant{},
		}
		addService := NewAddService(repo)
		newPlant := plant.Plant{
			Name: "test",
		}
		result, err := addService.Add(uid, newPlant)

		require.NoError(t, err)
		require.NotNil(t, result)
	})
}
