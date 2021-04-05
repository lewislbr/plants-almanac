package add

import (
	"time"

	"plants/plant"

	"github.com/google/uuid"
)

type (
	Inserter interface {
		InsertOne(string, plant.Plant) (interface{}, error)
	}

	service struct {
		r Inserter
	}
)

func NewService(r Inserter) *service {
	return &service{r}
}

func (s *service) Add(uid string, new plant.Plant) error {
	if new.Name == "" {
		return plant.ErrMissingData
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()
	new.EditedAt = time.Now().UTC()

	_, err := s.r.InsertOne(uid, new)
	if err != nil {
		return err
	}

	return nil
}
