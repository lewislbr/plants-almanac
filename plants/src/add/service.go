package add

import (
	"time"

	p "plants/src/plant"

	"github.com/google/uuid"
)

// Service provides plant add operations
type Service interface {
	AddPlant(string, p.Plant) interface{}
}

// Repository provides access to the plant storage
type Repository interface {
	InsertOne(string, p.Plant) interface{}
}

type service struct {
	r Repository
}

// AddPlant adds a plant
func (s *service) AddPlant(uid string, plant p.Plant) interface{} {
	plant.ID = p.ID(uuid.New().String())
	plant.CreatedAt = time.Now().UTC()

	return s.r.InsertOne(uid, plant)
}

// NewService creates an add service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
