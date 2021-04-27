package list

import (
	"plants/plant"
)

type (
	Finder interface {
		FindAll(string) ([]plant.Plant, error)
		FindOne(string, string) (plant.Plant, error)
	}

	service struct {
		repo Finder
	}
)

func NewService(repo Finder) *service {
	return &service{repo}
}

func (s *service) ListAll(uid string) ([]plant.Plant, error) {
	result, err := s.repo.FindAll(uid)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) ListOne(uid, id string) (plant.Plant, error) {
	if id == "" {
		return plant.Plant{}, plant.ErrMissingData
	}

	result, err := s.repo.FindOne(uid, id)
	if err != nil {
		return plant.Plant{}, plant.ErrNotFound
	}

	return result, nil
}
