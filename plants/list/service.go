package list

import (
	p "plants/plant"
)

type listService struct {
	r p.Repository
}

func NewListService(r p.Repository) *listService {
	return &listService{r}
}

func (s *listService) ListAll(uid string) ([]p.Plant, error) {
	result, err := s.r.FindAll(uid)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *listService) ListOne(uid string, id string) (p.Plant, error) {
	if id == "" {
		return p.Plant{}, p.ErrMissingData
	}

	result, err := s.r.FindOne(uid, id)
	if err != nil {
		return p.Plant{}, p.ErrNotFound
	}

	return result, nil
}
