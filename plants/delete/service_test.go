package delete

import (
	"testing"

	"plants/plant"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const uid = "123"

func TestCreate(t *testing.T) {
	t.Run("should error when there are missing required fields", func(t *testing.T) {
		t.Parallel()

		d := &MockDeleter{}

		d.On("DeleteOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(1), nil)

		ds := NewService(d)
		id := ""
		err := ds.Delete(uid, id)

		require.EqualError(t, err, plant.ErrMissingData.Error())
	})

	t.Run("should error when there are no matches", func(t *testing.T) {
		t.Parallel()

		d := &MockDeleter{}

		d.On("DeleteOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(0), nil)

		ds := NewService(d)
		id := "124"
		err := ds.Delete(uid, id)

		require.EqualError(t, err, plant.ErrNotFound.Error())
	})

	t.Run("should delete a plant with no error", func(t *testing.T) {
		t.Parallel()

		d := &MockDeleter{}

		d.On("DeleteOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(1), nil)

		ds := NewService(d)
		id := "123"
		err := ds.Delete(uid, id)

		require.NoError(t, err)
	})
}
