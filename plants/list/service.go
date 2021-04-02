package list

import (
	"plants/plant"
	"plants/storage"
)

type service struct {
	r storage.Repository
}

func NewService(r storage.Repository) *service {
	return &service{r}
}

func (s *service) ListAll(uid string) ([]plant.Plant, error) {
	result, err := s.r.FindAll(uid)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) ListOne(uid, id string) (plant.Plant, error) {
	if id == "" {
		return plant.Plant{}, plant.ErrMissingData
	}

	result, err := s.r.FindOne(uid, id)
	if err != nil {
		return plant.Plant{}, plant.ErrNotFound
	}

	return result, nil
}
