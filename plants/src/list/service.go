package list

import (
	p "plants/src/plant"
	"plants/src/storage"

	"github.com/pkg/errors"
)

// ListAll lists all plants.
func ListAll(uid string) ([]p.Plant, error) {
	result, err := storage.FindAll(uid)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}

// ListOne lists a plant.
func ListOne(uid string, id string) (p.Plant, error) {
	if id == "" {
		return p.Plant{}, p.ErrMissingData
	}

	result, err := storage.FindOne(uid, id)
	if err != nil {
		return p.Plant{}, p.ErrNotFound
	}

	return result, nil
}
