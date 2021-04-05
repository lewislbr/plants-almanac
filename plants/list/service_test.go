package list

import (
	"errors"
	"testing"

	"plants/plant"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const uid = "123"

func TestCreate(t *testing.T) {
	t.Run("should list all plants with no error", func(t *testing.T) {
		t.Parallel()

		f := &MockFinder{}

		f.On("FindAll", mock.AnythingOfType("string")).Return([]plant.Plant{}, nil)

		ls := NewService(f)
		result, err := ls.ListAll(uid)

		require.NoError(t, err)
		require.NotNil(t, result)
	})

	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		f := &MockFinder{}

		f.On("FindAll", mock.AnythingOfType("string")).Return([]plant.Plant{}, nil)

		ls := NewService(f)
		id := ""
		result, err := ls.ListOne(uid, id)

		require.EqualError(t, err, plant.ErrMissingData.Error())
		require.Equal(t, plant.Plant{}, result)
	})

	t.Run("should error when a plant is not found", func(t *testing.T) {
		t.Parallel()

		f := &MockFinder{}

		f.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, errors.New("not found"))

		ls := NewService(f)
		id := "123"
		result, err := ls.ListOne(uid, id)

		require.EqualError(t, err, plant.ErrNotFound.Error())
		require.Equal(t, plant.Plant{}, result)
	})

	t.Run("should list a plant with no error", func(t *testing.T) {
		t.Parallel()

		f := &MockFinder{}

		f.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, nil)

		ls := NewService(f)
		id := "123"
		result, err := ls.ListOne(uid, id)

		require.NoError(t, err)
		require.NotNil(t, result)
	})
}
