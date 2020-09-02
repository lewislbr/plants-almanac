package edit

import "plants/pkg/entity"

// Service provides item edit operations
type Service interface {
	EditPlant(string, entity.Plant) int64
}

// Repository provides access to the item storage
type Repository interface {
	EditOne(string, entity.Plant) int64
}

type service struct {
	r Repository
}

// NewService creates an edit service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// EditPlant edits a plant
func (s *service) EditPlant(id string, new entity.Plant) int64 {
	return s.r.EditOne(id, new)
}
