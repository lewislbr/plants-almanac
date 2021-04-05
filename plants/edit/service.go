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
		ls Lister
		r  Updater
	}
)

func NewService(ls Lister, r Updater) *service {
	return &service{ls, r}
}

func (s *service) Edit(uid, id string, update plant.Plant) error {
	if update.Name == "" {
		return plant.ErrMissingData
	}

	exist, err := s.ls.ListOne(uid, id)
	if err != nil {
		return plant.ErrNotFound
	}

	update.CreatedAt = exist.CreatedAt
	update.EditedAt = time.Now().UTC()

	result, err := s.r.UpdateOne(uid, id, update)
	if err != nil {
		return err
	}
	if result == 0 {
		return plant.ErrNotFound
	}

	return nil
}
