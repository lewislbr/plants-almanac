package delete

import (
	"plants/plant"
)

type (
	Deleter interface {
		DeleteOne(string, string) (int64, error)
	}

	service struct {
		repo Deleter
	}
)

func NewService(repo Deleter) *service {
	return &service{repo}
}

func (s *service) Delete(uid, id string) error {
	if id == "" {
		return plant.ErrMissingData
	}

	result, err := s.repo.DeleteOne(uid, id)
	if err != nil {
		return err
	}
	if result == 0 {
		return plant.ErrNotFound
	}

	return nil
}
