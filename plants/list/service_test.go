package list

import (
	"errors"
	"testing"

	"plants/plant"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const userID = "123"

func TestCreate(t *testing.T) {
	t.Run("should list all plants with no error", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindAll", mock.AnythingOfType("string")).Return([]plant.Plant{}, nil)

		listSvc := NewService(repo)
		result, err := listSvc.ListAll(userID)

		require.NoError(t, err)
		require.NotNil(t, result)
	})

	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindAll", mock.AnythingOfType("string")).Return([]plant.Plant{}, nil)

		listSvc := NewService(repo)
		plantID := ""
		result, err := listSvc.ListOne(userID, plantID)

		require.EqualError(t, err, plant.ErrMissingData.Error())
		require.Equal(t, plant.Plant{}, result)
	})

	t.Run("should error when a plant is not found", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, errors.New("not found"))

		listSvc := NewService(repo)
		plantID := "123"
		result, err := listSvc.ListOne(userID, plantID)

		require.EqualError(t, err, plant.ErrNotFound.Error())
		require.Equal(t, plant.Plant{}, result)
	})

	t.Run("should list a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, nil)

		listSvc := NewService(repo)
		plantID := "123"
		result, err := listSvc.ListOne(userID, plantID)

		require.NoError(t, err)
		require.NotNil(t, result)
	})
}
