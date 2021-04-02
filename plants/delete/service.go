package delete

import (
	"plants/plant"
	"plants/storage"
)

type DeleteService interface {
	Delete(string, string) error
}

type deleteService struct {
	r storage.Repository
}

func NewDeleteService(r storage.Repository) *deleteService {
	return &deleteService{r}
}

func (ds *deleteService) Delete(uid, id string) error {
	if id == "" {
		return plant.ErrMissingData
	}

	result, err := ds.r.DeleteOne(uid, id)
	if err != nil {
		return err
	}
	if result == 0 {
		return plant.ErrNotFound
	}

	return nil
}
