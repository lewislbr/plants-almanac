package delete

// Service provides item delete operations
type Service interface {
	DeletePlant(string) int64
}

// Repository provides access to the item storage
type Repository interface {
	DeleteOne(string) int64
}

type service struct {
	r Repository
}

// NewService creates a delete service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// DeletePlant deletes a plant
func (s *service) DeletePlant(id string) int64 {
	return s.r.DeleteOne(id)
}
