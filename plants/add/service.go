package add

import (
	"time"

	"plants/plant"

	"github.com/google/uuid"
)

type (
	Inserter interface {
		InsertOne(string, plant.Plant) (interface{}, error)
	}

	service struct {
		repo Inserter
	}
)

func NewService(repo Inserter) *service {
	return &service{repo}
}

func (s *service) Add(uid string, new plant.Plant) error {
	if new.Name == "" {
		return plant.ErrMissingData
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()
	new.EditedAt = time.Now().UTC()

	_, err := s.repo.InsertOne(uid, new)

	return err
}
