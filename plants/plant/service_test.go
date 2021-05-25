package plant

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const userID = "123"

func TestAdd(t *testing.T) {
	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("Insert", mock.AnythingOfType("string"), mock.AnythingOfType("Plant")).Return("", nil)

		service := NewService(repo)
		newPlant := Plant{
			Name: "",
		}
		err := service.Add(userID, newPlant)

		require.EqualError(t, err, ErrMissingData.Error())
	})

	t.Run("should create a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("Insert", mock.AnythingOfType("string"), mock.AnythingOfType("Plant")).Return("", nil)

		service := NewService(repo)
		newPlant := Plant{
			Name: "test",
		}
		err := service.Add(userID, newPlant)

		require.NoError(t, err)
	})
}

func TestList(t *testing.T) {
	t.Run("should list all plants with no error", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindAll", mock.AnythingOfType("string")).Return([]Plant{}, nil)

		service := NewService(repo)
		result, err := service.ListAll(userID)

		require.NoError(t, err)
		require.NotNil(t, result)
	})

	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindAll", mock.AnythingOfType("string")).Return([]Plant{}, nil)

		service := NewService(repo)
		plantID := ""
		result, err := service.ListOne(userID, plantID)

		require.EqualError(t, err, ErrMissingData.Error())
		require.Equal(t, Plant{}, result)
	})

	t.Run("should error when a plant is not found", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(Plant{}, errors.New("not found"))

		service := NewService(repo)
		plantID := "123"
		result, err := service.ListOne(userID, plantID)

		require.EqualError(t, err, ErrNotFound.Error())
		require.Equal(t, Plant{}, result)
	})

	t.Run("should list a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(Plant{}, nil)

		service := NewService(repo)
		plantID := "123"
		result, err := service.ListOne(userID, plantID)

		require.NoError(t, err)
		require.NotNil(t, result)
	})
}

func TestEdit(t *testing.T) {
	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(Plant{}, nil)
		repo.On("Update", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("Plant")).Return(int64(1), nil)

		service := NewService(repo)
		plantID := "123"
		newPlant := Plant{
			Name: "",
		}
		err := service.Edit(userID, plantID, newPlant)

		require.EqualError(t, err, ErrMissingData.Error())
	})

	t.Run("should error when the plant is not found", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(Plant{}, errors.New("not found"))
		repo.On("Update", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("Plant")).Return(int64(0), nil)

		service := NewService(repo)
		plantID := "123"
		newPlant := Plant{
			Name: "test",
		}
		err := service.Edit(userID, plantID, newPlant)

		require.EqualError(t, err, ErrNotFound.Error())
	})

	t.Run("should edit a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(Plant{}, nil)
		repo.On("Update", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("Plant")).Return(int64(1), nil)

		service := NewService(repo)
		plantID := "123"
		newPlant := Plant{
			Name: "test",
		}
		err := service.Edit(userID, plantID, newPlant)

		require.NoError(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("Delete", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(1), nil)

		service := NewService(repo)
		plantID := ""
		err := service.Delete(userID, plantID)

		require.EqualError(t, err, ErrMissingData.Error())
	})

	t.Run("should error when there are no matches", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("Delete", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(0), nil)

		service := NewService(repo)
		plantID := "124"
		err := service.Delete(userID, plantID)

		require.EqualError(t, err, ErrNotFound.Error())
	})

	t.Run("should delete a plant with no error", func(t *testing.T) {
		t.Parallel()

		repo := &mockRepository{}

		repo.On("Delete", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(1), nil)

		service := NewService(repo)
		plantID := "123"
		err := service.Delete(userID, plantID)

		require.NoError(t, err)
	})
}
