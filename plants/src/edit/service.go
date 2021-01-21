package edit

import (
	"time"

	"plants/src/list"
	p "plants/src/plant"
	"plants/src/storage"

	"github.com/pkg/errors"
)

// Edit edits a plant.
func Edit(uid string, id string, update p.Plant) (int64, error) {
	if update.Name == "" {
		return 0, p.ErrMissingData
	}

	exist, err := list.ListOne(uid, id)
	if err != nil {
		return 0, p.ErrNotFound
	}

	update.CreatedAt = exist.CreatedAt
	update.EditedAt = time.Now().UTC()

	result, err := storage.UpdateOne(uid, id, update)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result, nil
}
