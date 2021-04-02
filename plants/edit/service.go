package edit

import (
	"time"

	"plants/list"
	"plants/plant"
	"plants/storage"
)

type EditService interface {
	Edit(string, string, plant.Plant) error
}

type editService struct {
	ls list.ListService
	r  storage.Repository
}

func NewEditService(ls list.ListService, r storage.Repository) *editService {
	return &editService{ls, r}
}

func (es *editService) Edit(uid, id string, update plant.Plant) error {
	if update.Name == "" {
		return plant.ErrMissingData
	}

	exist, err := es.ls.ListOne(uid, id)
	if err != nil {
		return plant.ErrNotFound
	}

	update.CreatedAt = exist.CreatedAt
	update.EditedAt = time.Now().UTC()

	result, err := es.r.UpdateOne(uid, id, update)
	if err != nil {
		return err
	}
	if result == 0 {
		return plant.ErrNotFound
	}

	return nil
}
