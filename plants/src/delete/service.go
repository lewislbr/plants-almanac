package delete

import (
	p "plants/src/plant"
	"plants/src/storage"

	"github.com/pkg/errors"
)

// Delete deletes a plant.
func Delete(uid string, id string) (int64, error) {
	if id == "" {
		return 0, p.ErrMissingData
	}

	result, err := storage.DeleteOne(uid, id)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result, nil
}
