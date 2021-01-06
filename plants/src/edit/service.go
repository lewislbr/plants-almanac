package edit

import (
	"time"

	p "plants/src/plant"

	"github.com/pkg/errors"
)

// Service defines a service to edit a plant.
type Service interface {
	Edit(string, p.ID, p.Plant) (int64, error)
}

type repository interface {
	UpdateOne(string, p.ID, p.Plant) (int64, error)
}

type service struct {
	r repository
}

// NewService creates an edit service with the necessary dependencies.
func NewService(r repository) Service {
	return &service{r}
}

// Edit edits a plant.
func (s *service) Edit(uid string, id p.ID, update p.Plant) (int64, error) {
	update.EditedAt = time.Now().UTC()

	result, err := s.r.UpdateOne(uid, id, update)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result, nil
}
