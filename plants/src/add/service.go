package add

import (
	"time"

	p "plants/src/plant"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Service defines a service to add a plant.
type Service interface {
	Add(string, p.Plant) (interface{}, error)
}

type repository interface {
	InsertOne(string, p.Plant) (interface{}, error)
}

type service struct {
	r repository
}

// NewService creates an add service with the necessary dependencies.
func NewService(r repository) Service {
	return &service{r}
}

// Add adds a plant.
func (s *service) Add(uid string, new p.Plant) (interface{}, error) {
	new.ID = p.ID(uuid.New().String())
	new.CreatedAt = time.Now().UTC()

	result, err := s.r.InsertOne(uid, new)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
