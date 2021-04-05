package edit

import (
	"errors"
	"testing"

	"plants/list"
	"plants/plant"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const uid = "123"

func TestCreate(t *testing.T) {
	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		f := &list.MockFinder{}
		u := &MockUpdater{}

		f.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, nil)
		u.On("UpdateOne", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(int64(1), nil)

		ls := list.NewService(f)
		es := NewService(ls, u)
		id := "123"
		newPlant := plant.Plant{
			Name: "",
		}
		err := es.Edit(uid, id, newPlant)

		require.EqualError(t, err, plant.ErrMissingData.Error())
	})

	t.Run("should error when the plant is not found", func(t *testing.T) {
		t.Parallel()

		f := &list.MockFinder{}
		u := &MockUpdater{}

		f.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, errors.New("not found"))
		u.On("UpdateOne", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(int64(0), nil)

		ls := list.NewService(f)
		es := NewService(ls, u)
		id := "123"
		newPlant := plant.Plant{
			Name: "test",
		}
		err := es.Edit(uid, id, newPlant)

		require.EqualError(t, err, plant.ErrNotFound.Error())
	})

	t.Run("should edit a plant with no error", func(t *testing.T) {
		t.Parallel()

		f := &list.MockFinder{}
		u := &MockUpdater{}

		f.On("FindOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, nil)
		u.On("UpdateOne", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(int64(1), nil)

		ls := list.NewService(f)
		es := NewService(ls, u)
		id := "123"
		newPlant := plant.Plant{
			Name: "test",
		}
		err := es.Edit(uid, id, newPlant)

		require.NoError(t, err)
	})
}
