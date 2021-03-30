package add

import (
	"time"

	"plants/plant"

	"github.com/google/uuid"
)

type addService struct {
	r plant.Repository
}

func NewAddService(r plant.Repository) *addService {
	return &addService{r}
}

func (as *addService) Add(uid string, new plant.Plant) (interface{}, error) {
	if new.Name == "" {
		return nil, plant.ErrMissingData
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()
	new.EditedAt = time.Now().UTC()

	result, err := as.r.InsertOne(uid, new)
	if err != nil {
		return nil, err
	}

	return result, nil
}
