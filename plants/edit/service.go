package edit

import (
	"time"

	"plants/plant"
)

type (
	repository interface {
		FindOne(string, string) (plant.Plant, error)
		UpdateOne(string, string, plant.Plant) (int64, error)
	}

	service struct {
		repo repository
	}
)

func NewService(repo repository) *service {
	return &service{repo}
}

func (s *service) Edit(userID, plantID string, update plant.Plant) error {
	if update.Name == "" {
		return plant.ErrMissingData
	}

	exist, err := s.repo.FindOne(userID, plantID)
	if err != nil {
		return plant.ErrNotFound
	}

	update.CreatedAt = exist.CreatedAt
	update.EditedAt = time.Now().UTC()

	result, err := s.repo.UpdateOne(userID, plantID, update)
	if err != nil {
		return err
	}
	if result == 0 {
		return plant.ErrNotFound
	}

	return nil
}
