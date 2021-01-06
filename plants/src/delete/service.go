package delete

import (
	p "plants/src/plant"

	"github.com/pkg/errors"
)

// Service defines a service to delete a plant.
type Service interface {
	Delete(string, p.ID) (int64, error)
}

type repository interface {
	DeleteOne(string, p.ID) (int64, error)
}

type service struct {
	r repository
}

// NewService creates a delete service with the necessary dependencies.
func NewService(r repository) Service {
	return &service{r}
}

// Delete deletes a plant.
func (s *service) Delete(uid string, id p.ID) (int64, error) {
	result, err := s.r.DeleteOne(uid, id)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result, nil
}
