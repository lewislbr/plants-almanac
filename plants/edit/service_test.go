package edit

import (
	"errors"
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

		repo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, nil)
		repo.On("UpdateOne", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(int64(1), nil)

		editSvc := NewService(repo)
		plantID := "123"
		newPlant := plant.Plant{
			Name: "",
		}
		err := editSvc.Edit(userID, plantID, newPlant)

		require.EqualError(t, err, plant.ErrMissingData.Error())
	})

	t.Run("should error when the plant is not found", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, errors.New("not found"))
		repo.On("UpdateOne", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(int64(0), nil)

		editSvc := NewService(repo)
		plantID := "123"
		newPlant := plant.Plant{
			Name: "test",
		}
		err := editSvc.Edit(userID, plantID, newPlant)

		require.EqualError(t, err, plant.ErrNotFound.Error())
	})

	t.Run("should edit a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, nil)
		repo.On("UpdateOne", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(int64(1), nil)

		editSvc := NewService(repo)
		plantID := "123"
		newPlant := plant.Plant{
			Name: "test",
		}
		err := editSvc.Edit(userID, plantID, newPlant)

		require.NoError(t, err)
	})
}
