package list

import (
	"plants/plant"
	"plants/storage"
)

type ListService interface {
	ListAll(string) ([]plant.Plant, error)
	ListOne(string, string) (plant.Plant, error)
}

type listService struct {
	r storage.Repository
}

func NewListService(r storage.Repository) *listService {
	return &listService{r}
}

func (s *listService) ListAll(uid string) ([]plant.Plant, error) {
	result, err := s.r.FindAll(uid)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *listService) ListOne(uid, id string) (plant.Plant, error) {
	if id == "" {
		return plant.Plant{}, plant.ErrMissingData
	}

	result, err := s.r.FindOne(uid, id)
	if err != nil {
		return plant.Plant{}, plant.ErrNotFound
	}

	return result, nil
}
