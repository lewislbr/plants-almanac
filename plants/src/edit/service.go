package edit

import (
	"time"

	p "plants/src/plant"
)

// Service provides plant edit operations
type Service interface {
	EditPlant(p.ID, p.Plant) int64
}

// Repository provides access to the plant storage
type Repository interface {
	UpdateOne(p.ID, p.Plant) int64
}

type service struct {
	r Repository
}

// EditPlant edits a plant
func (s *service) EditPlant(id p.ID, plant p.Plant) int64 {
	plant.EditedAt = time.Now().UTC()

	return s.r.UpdateOne(id, plant)
}

// NewService creates an edit service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
