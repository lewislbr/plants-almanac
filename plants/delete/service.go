package delete

import (
	p "plants/plant"

	"github.com/pkg/errors"
)

type deleteService struct {
	r p.Repository
}

// NewDeleteService initializes a create service with the necessary dependencies.
func NewDeleteService(r p.Repository) *deleteService {
	return &deleteService{r}
}

// Delete deletes a plant.
func (ds *deleteService) Delete(uid string, id string) (int64, error) {
	if id == "" {
		return 0, p.ErrMissingData
	}

	result, err := ds.r.DeleteOne(uid, id)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result, nil
}
