package add

import "plants/pkg/entity"

// Service provides item add operations
type Service interface {
	AddPlant(entity.Plant) interface{}
}

// Repository provides access to the item storage
type Repository interface {
	InsertOne(entity.Plant) interface{}
}

type service struct {
	r Repository
}

// NewService creates an add service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddPlant adds a plant
func (s *service) AddPlant(plant entity.Plant) interface{} {
	return s.r.InsertOne(plant)
}
