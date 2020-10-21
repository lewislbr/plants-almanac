package delete

import p "plants/pkg/plant"

// Service provides plant delete operations
type Service interface {
	DeletePlant(p.ID) int64
}

// Repository provides access to the plant storage
type Repository interface {
	DeleteOne(p.ID) int64
}

type service struct {
	r Repository
}

// DeletePlant deletes a plant
func (s *service) DeletePlant(id p.ID) int64 {
	return s.r.DeleteOne(id)
}

// NewService creates a delete service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
