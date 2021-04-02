package add

import (
	"time"

	"plants/plant"
	"plants/storage"

	"github.com/google/uuid"
)

type AddService interface {
	Add(string, plant.Plant) error
}

type addService struct {
	r storage.Repository
}

func NewAddService(r storage.Repository) *addService {
	return &addService{r}
}

func (as *addService) Add(uid string, new plant.Plant) error {
	if new.Name == "" {
		return plant.ErrMissingData
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()
	new.EditedAt = time.Now().UTC()

	_, err := as.r.InsertOne(uid, new)
	if err != nil {
		return err
	}

	return nil
}
