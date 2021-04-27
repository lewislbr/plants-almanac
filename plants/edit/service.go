package edit

import (
	"time"

	"plants/plant"
)

type (
	Updater interface {
		UpdateOne(string, string, plant.Plant) (int64, error)
	}

	Lister interface {
		ListAll(string) ([]plant.Plant, error)
		ListOne(string, string) (plant.Plant, error)
	}

	service struct {
		svc  Lister
		repo Updater
	}
)

func NewService(svc Lister, repo Updater) *service {
	return &service{svc, repo}
}

func (s *service) Edit(uid, id string, update plant.Plant) error {
	if update.Name == "" {
		return plant.ErrMissingData
	}

	exist, err := s.svc.ListOne(uid, id)
	if err != nil {
		return plant.ErrNotFound
	}

	update.CreatedAt = exist.CreatedAt
	update.EditedAt = time.Now().UTC()

	result, err := s.repo.UpdateOne(uid, id, update)
	if err != nil {
		return err
	}
	if result == 0 {
		return plant.ErrNotFound
	}

	return nil
}
