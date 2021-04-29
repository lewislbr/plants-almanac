package list

import (
	"plants/plant"
)

type (
	repository interface {
		FindAll(string) ([]plant.Plant, error)
		FindOne(string, string) (plant.Plant, error)
	}

	service struct {
		repo repository
	}
)

func NewService(repo repository) *service {
	return &service{repo}
}

func (s *service) ListAll(userID string) ([]plant.Plant, error) {
	result, err := s.repo.FindAll(userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) ListOne(userID, plantID string) (plant.Plant, error) {
	if plantID == "" {
		return plant.Plant{}, plant.ErrMissingData
	}

	result, err := s.repo.FindOne(userID, plantID)
	if err != nil {
		return plant.Plant{}, plant.ErrNotFound
	}

	return result, nil
}
