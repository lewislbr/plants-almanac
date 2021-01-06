package list

import (
	p "plants/src/plant"

	"github.com/pkg/errors"
)

// Service defines a service to list plants.
type Service interface {
	ListAll(string) ([]p.Plant, error)
	ListOne(string, p.ID) (p.Plant, error)
}

type repository interface {
	FindAll(string) ([]p.Plant, error)
	FindOne(string, p.ID) (p.Plant, error)
}

type service struct {
	r repository
}

// NewService creates a list service with the necessary dependencies.
func NewService(r repository) Service {
	return &service{r}
}

// ListAll lists all plants.
func (s *service) ListAll(uid string) ([]p.Plant, error) {
	result, err := s.r.FindAll(uid)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}

// ListOne lists a plant.
func (s *service) ListOne(uid string, id p.ID) (p.Plant, error) {
	result, err := s.r.FindOne(uid, id)
	if err != nil {
		return p.Plant{}, errors.Wrap(err, "")
	}

	return result, nil
}
