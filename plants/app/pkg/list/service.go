package list

import "plants/pkg/entity"

// Service provides item list operations
type Service interface {
	GetPlants() []*entity.Plant
	GetPlant(string) *entity.Plant
}

// Repository provides access to the item storage
type Repository interface {
	FindAll() []*entity.Plant
	FindOne(string) *entity.Plant
}

type service struct {
	r Repository
}

// NewService creates a list service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetPlants returns all plants
func (s *service) GetPlants() []*entity.Plant {
	return s.r.FindAll()
}

// GetPlant returns a plant
func (s *service) GetPlant(id string) *entity.Plant {
	return s.r.FindOne(id)
}
