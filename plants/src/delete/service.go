package delete

import p "plants/src/plant"

// Service provides plant delete operations
type Service interface {
	DeletePlant(string, p.ID) int64
}

// Repository provides access to the plant storage
type Repository interface {
	DeleteOne(string, p.ID) int64
}

type service struct {
	r Repository
}

// DeletePlant deletes a plant
func (s *service) DeletePlant(uid string, id p.ID) int64 {
	return s.r.DeleteOne(uid, id)
}

// NewService creates a delete service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
