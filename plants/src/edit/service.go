package edit

import (
	"time"

	p "plants/src/plant"

	"github.com/pkg/errors"
)

type editService struct {
	ls p.ListService
	r  p.Repository
}

// NewEditService initializes a create service with the necessary dependencies.
func NewEditService(ls p.ListService, r p.Repository) p.EditService {
	return editService{ls, r}
}

// Edit edits a plant.
func (s editService) Edit(uid string, id string, update p.Plant) (int64, error) {
	if update.Name == "" {
		return 0, p.ErrMissingData
	}

	exist, err := s.ls.ListOne(uid, id)
	if err != nil {
		return 0, p.ErrNotFound
	}

	update.CreatedAt = exist.CreatedAt
	update.EditedAt = time.Now().UTC()

	result, err := s.r.UpdateOne(uid, id, update)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result, nil
}
