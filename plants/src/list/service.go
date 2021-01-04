package list

import (
	p "plants/src/plant"

	"github.com/pkg/errors"
)

// Service provides plant list operations
type Service interface {
	ListPlants(string) ([]p.Plant, error)
	ListPlant(string, p.ID) (p.Plant, error)
}

// Repository provides access to the plant storage
type Repository interface {
	FindAll(string) ([]p.Plant, error)
	FindOne(string, p.ID) (p.Plant, error)
}

type service struct {
	r Repository
}

// ListPlants lists all plants
func (s *service) ListPlants(uid string) ([]p.Plant, error) {
	result, err := s.r.FindAll(uid)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}

// ListPlant lists a plant
func (s *service) ListPlant(uid string, id p.ID) (p.Plant, error) {
	result, err := s.r.FindOne(uid, id)
	if err != nil {
		return p.Plant{}, errors.Wrap(err, "")
	}

	return result, nil
}

// NewService creates a list service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
