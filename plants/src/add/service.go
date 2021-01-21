package add

import (
	"time"

	p "plants/src/plant"
	"plants/src/storage"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Add adds a plant.
func Add(uid string, new p.Plant) (interface{}, error) {
	if new.Name == "" {
		return nil, p.ErrMissingData
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()
	new.EditedAt = time.Now().UTC()

	result, err := storage.InsertOne(uid, new)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
