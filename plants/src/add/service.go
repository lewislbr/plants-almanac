package add

import (
	"time"

	p "plants/src/plant"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Service provides plant add operations
type Service interface {
	AddPlant(string, p.Plant) (interface{}, error)
}

// Repository provides access to the plant storage
type Repository interface {
	InsertOne(string, p.Plant) (interface{}, error)
}

type service struct {
	r Repository
}

// AddPlant adds a plant
func (s *service) AddPlant(uid string, plant p.Plant) (interface{}, error) {
	plant.ID = p.ID(uuid.New().String())
	plant.CreatedAt = time.Now().UTC()

	result, err := s.r.InsertOne(uid, plant)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}

// NewService creates an add service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
