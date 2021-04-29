package add

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

		repo.On("InsertOne", mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return("", nil)

		addSvc := NewService(repo)
		newPlant := plant.Plant{
			Name: "",
		}
		err := addSvc.Add(userID, newPlant)

		require.EqualError(t, err, plant.ErrMissingData.Error())
	})

	t.Run("should create a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("InsertOne", mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return("", nil)

		addSvc := NewService(repo)
		newPlant := plant.Plant{
			Name: "test",
		}
		err := addSvc.Add(userID, newPlant)

		require.NoError(t, err)
	})
}
