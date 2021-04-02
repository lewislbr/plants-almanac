package delete

import (
	"plants/plant"
	"plants/storage"
)

type service struct {
	r storage.Repository
}

func NewService(r storage.Repository) *service {
	return &service{r}
}

func (s *service) Delete(uid, id string) error {
	if id == "" {
		return plant.ErrMissingData
	}

	result, err := s.r.DeleteOne(uid, id)
	if err != nil {
		return err
	}
	if result == 0 {
		return plant.ErrNotFound
	}

	return nil
}
