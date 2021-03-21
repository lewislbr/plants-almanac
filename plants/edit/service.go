package edit

import (
	"time"

	p "plants/plant"
)

type editService struct {
	ls p.ListService
	r  p.Repository
}

func NewEditService(ls p.ListService, r p.Repository) *editService {
	return &editService{ls, r}
}

func (es *editService) Edit(uid string, id string, update p.Plant) (int64, error) {
	if update.Name == "" {
		return 0, p.ErrMissingData
	}

	exist, err := es.ls.ListOne(uid, id)
	if err != nil {
		return 0, p.ErrNotFound
	}

	update.CreatedAt = exist.CreatedAt
	update.EditedAt = time.Now().UTC()

	result, err := es.r.UpdateOne(uid, id, update)
	if err != nil {
		return 0, err
	}

	return result, nil
}
