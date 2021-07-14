package plant

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const userID = "123"

func TestAdd(t *testing.T) {
	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		plantRepo := &mockPlantRepo{}

		plantRepo.On("Insert", mock.AnythingOfType("string"), mock.AnythingOfType("Plant")).Return("", nil)

		plantService := NewService(plantRepo)
		newPlant := Plant{
			Name: "",
		}
		err := plantService.Add(userID, newPlant)

		require.ErrorIs(t, err, ErrMissingData)
	})

	t.Run("should create a plant with no error", func(t *testing.T) {
		t.Parallel()

		plantRepo := &mockPlantRepo{}

		plantRepo.On("Insert", mock.AnythingOfType("string"), mock.AnythingOfType("Plant")).Return("", nil)

		plantService := NewService(plantRepo)
		newPlant := Plant{
			Name: "test",
		}
		err := plantService.Add(userID, newPlant)

		require.NoError(t, err)
	})
}

func TestList(t *testing.T) {
	t.Run("should list all plants with no error", func(t *testing.T) {
		t.Parallel()

		plantRepo := &mockPlantRepo{}

		plantRepo.On("FindAll", mock.AnythingOfType("string")).Return([]Plant{}, nil)

		plantService := NewService(plantRepo)
		result, err := plantService.ListAll(userID)

		require.NoError(t, err)
		require.NotNil(t, result)
	})

	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		plantRepo := &mockPlantRepo{}

		plantRepo.On("FindAll", mock.AnythingOfType("string")).Return([]Plant{}, nil)

		plantService := NewService(plantRepo)
		plantID := ""
		result, err := plantService.ListOne(userID, plantID)

		require.ErrorIs(t, err, ErrMissingData)
		require.Equal(t, Plant{}, result)
	})

	t.Run("should error when a plant is not found", func(t *testing.T) {
		t.Parallel()

		plantRepo := &mockPlantRepo{}

		plantRepo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(Plant{}, ErrNotFound)

		plantService := NewService(plantRepo)
		plantID := "123"
		result, err := plantService.ListOne(userID, plantID)

		require.ErrorIs(t, err, ErrNotFound)
		require.Equal(t, Plant{}, result)
	})

	t.Run("should list a plant with no error", func(t *testing.T) {
		t.Parallel()

		plantRepo := &mockPlantRepo{}

		plantRepo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(Plant{}, nil)

		plantService := NewService(plantRepo)
		plantID := "123"
		result, err := plantService.ListOne(userID, plantID)

		require.NoError(t, err)
		require.NotNil(t, result)
	})
}

func TestEdit(t *testing.T) {
	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		plantRepo := &mockPlantRepo{}

		plantRepo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(Plant{}, nil)
		plantRepo.On("Update", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("Plant")).Return(int64(1), nil)

		plantService := NewService(plantRepo)
		plantID := "123"
		newPlant := Plant{
			Name: "",
		}
		err := plantService.Edit(userID, plantID, newPlant)

		require.ErrorIs(t, err, ErrMissingData)
	})

	t.Run("should error when the plant is not found", func(t *testing.T) {
		t.Parallel()

		plantRepo := &mockPlantRepo{}

		plantRepo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(Plant{}, ErrNotFound)
		plantRepo.On("Update", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("Plant")).Return(int64(0), nil)

		plantService := NewService(plantRepo)
		plantID := "123"
		newPlant := Plant{
			Name: "test",
		}
		err := plantService.Edit(userID, plantID, newPlant)

		require.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("should edit a plant with no error", func(t *testing.T) {
		t.Parallel()

		plantRepo := &mockPlantRepo{}

		plantRepo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(Plant{}, nil)
		plantRepo.On("Update", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("Plant")).Return(int64(1), nil)

		plantService := NewService(plantRepo)
		plantID := "123"
		newPlant := Plant{
			Name: "test",
		}
		err := plantService.Edit(userID, plantID, newPlant)

		require.NoError(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		plantRepo := &mockPlantRepo{}

		plantRepo.On("Delete", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(1), nil)

		plantService := NewService(plantRepo)
		plantID := ""
		err := plantService.Delete(userID, plantID)

		require.ErrorIs(t, err, ErrMissingData)
	})

	t.Run("should error when there are no matches", func(t *testing.T) {
		t.Parallel()

		plantRepo := &mockPlantRepo{}

		plantRepo.On("Delete", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(0), nil)

		plantService := NewService(plantRepo)
		plantID := "124"
		err := plantService.Delete(userID, plantID)

		require.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("should delete a plant with no error", func(t *testing.T) {
		t.Parallel()

		plantRepo := &mockPlantRepo{}

		plantRepo.On("Delete", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(1), nil)

		plantService := NewService(plantRepo)
		plantID := "123"
		err := plantService.Delete(userID, plantID)

		require.NoError(t, err)
	})
}
