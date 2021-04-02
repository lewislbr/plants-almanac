package add

import (
	"time"

	"plants/plant"
	"plants/storage"

	"github.com/google/uuid"
)

type service struct {
	r storage.Repository
}

func NewService(r storage.Repository) *service {
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
