package delete

import (
	"plants/plant"
)

type (
	Deleter interface {
		DeleteOne(string, string) (int64, error)
	}

	service struct {
		r Deleter
	}
)

func NewService(r Deleter) *service {
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
