package add

import (
	"time"

	p "plants/src/plant"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type addService struct {
	r p.Repository
}

// NewAddService initializes a create service with the necessary dependencies.
func NewAddService(r p.Repository) p.AddService {
	return addService{r}
}

// Add adds a plant.
func (s addService) Add(uid string, new p.Plant) (interface{}, error) {
	if new.Name == "" {
		return nil, p.ErrMissingData
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()
	new.EditedAt = time.Now().UTC()

	result, err := s.r.InsertOne(uid, new)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
