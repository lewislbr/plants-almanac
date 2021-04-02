package delete

import "plants/plant"

type deleteService struct {
	r plant.Repository
}

func NewDeleteService(r plant.Repository) *deleteService {
	return &deleteService{r}
}

func (ds *deleteService) Delete(uid, id string) (int64, error) {
	if id == "" {
		return 0, plant.ErrMissingData
	}

	result, err := ds.r.DeleteOne(uid, id)
	if err != nil {
		return 0, err
	}

	return result, nil
}
