package edit

import (
	"time"

	p "plants/src/plant"

	"github.com/pkg/errors"
)

// Service provides plant edit operations
type Service interface {
	EditPlant(string, p.ID, p.Plant) (int64, error)
}

// Repository provides access to the plant storage
type Repository interface {
	UpdateOne(string, p.ID, p.Plant) (int64, error)
}

type service struct {
	r Repository
}

// EditPlant edits a plant
func (s *service) EditPlant(uid string, id p.ID, plant p.Plant) (int64, error) {
	plant.EditedAt = time.Now().UTC()

	result, err := s.r.UpdateOne(uid, id, plant)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result, nil
}

// NewService creates an edit service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
