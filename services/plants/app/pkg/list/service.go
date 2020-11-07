package list

import p "plants/pkg/plant"

// Service provides plant list operations
type Service interface {
	ListPlants() []*p.Plant
	ListPlant(p.ID) *p.Plant
}

// Repository provides access to the plant storage
type Repository interface {
	FindAll() []*p.Plant
	FindOne(p.ID) *p.Plant
}

type service struct {
	r Repository
}

// ListPlants lists all plants
func (s *service) ListPlants() []*p.Plant {
	return s.r.FindAll()
}

// ListPlant lists a plant
func (s *service) ListPlant(id p.ID) *p.Plant {
	return s.r.FindOne(id)
}

// NewService creates a list service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
