package add

import (
	"time"

	"plants/plant"

	"github.com/google/uuid"
)

type (
	repository interface {
		InsertOne(string, plant.Plant) (interface{}, error)
	}

	service struct {
		repo repository
	}
)

func NewService(repo repository) *service {
	return &service{repo}
}

func (s *service) Add(userID string, new plant.Plant) error {
	if new.Name == "" {
		return plant.ErrMissingData
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()
	new.EditedAt = time.Now().UTC()

	_, err := s.repo.InsertOne(userID, new)

	return err
}
