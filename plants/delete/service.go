package delete

import (
	p "plants/plant"
)

type deleteService struct {
	r p.Repository
}

func NewDeleteService(r p.Repository) *deleteService {
	return &deleteService{r}
}

func (ds *deleteService) Delete(uid string, id string) (int64, error) {
	if id == "" {
		return 0, p.ErrMissingData
	}

	result, err := ds.r.DeleteOne(uid, id)
	if err != nil {
		return 0, err
	}

	return result, nil
}
