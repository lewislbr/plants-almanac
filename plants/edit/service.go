package edit

import (
	"time"

	"plants/plant"
)

type editService struct {
	ls plant.ListService
	r  plant.Repository
}

func NewEditService(ls plant.ListService, r plant.Repository) *editService {
	return &editService{ls, r}
}

func (es *editService) Edit(uid string, id string, update plant.Plant) (int64, error) {
	if update.Name == "" {
		return 0, plant.ErrMissingData
	}

	exist, err := es.ls.ListOne(uid, id)
	if err != nil {
		return 0, plant.ErrNotFound
	}

	update.CreatedAt = exist.CreatedAt
	update.EditedAt = time.Now().UTC()

	result, err := es.r.UpdateOne(uid, id, update)
	if err != nil {
		return 0, err
	}

	return result, nil
}
