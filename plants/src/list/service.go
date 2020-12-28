package list

import p "plants/src/plant"

// Service provides plant list operations
type Service interface {
	ListPlants(string) []*p.Plant
	ListPlant(string, p.ID) *p.Plant
}

// Repository provides access to the plant storage
type Repository interface {
	FindAll(string) []*p.Plant
	FindOne(string, p.ID) *p.Plant
}

type service struct {
	r Repository
}

// ListPlants lists all plants
func (s *service) ListPlants(uid string) []*p.Plant {
	return s.r.FindAll(uid)
}

// ListPlant lists a plant
func (s *service) ListPlant(uid string, id p.ID) *p.Plant {
	return s.r.FindOne(uid, id)
}

// NewService creates a list service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
