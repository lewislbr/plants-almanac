package delete

import (
	"plants/plant"
)

type (
	repository interface {
		DeleteOne(string, string) (int64, error)
	}

	service struct {
		repo repository
	}
)

func NewService(repo repository) *service {
	return &service{repo}
}

func (s *service) Delete(userID, plantID string) error {
	if plantID == "" {
		return plant.ErrMissingData
	}

	result, err := s.repo.DeleteOne(userID, plantID)
	if err != nil {
		return err
	}
	if result == 0 {
		return plant.ErrNotFound
	}

	return nil
}
