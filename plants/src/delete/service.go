package delete

import (
	p "plants/src/plant"

	"github.com/pkg/errors"
)

// Service provides plant delete operations
type Service interface {
	DeletePlant(string, p.ID) (int64, error)
}

// Repository provides access to the plant storage
type Repository interface {
	DeleteOne(string, p.ID) (int64, error)
}

type service struct {
	r Repository
}

// DeletePlant deletes a plant
func (s *service) DeletePlant(uid string, id p.ID) (int64, error) {
	result, err := s.r.DeleteOne(uid, id)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result, nil
}

// NewService creates a delete service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
