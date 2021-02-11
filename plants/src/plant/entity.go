package plant

import (
	"time"
)

// Plant defines the properties of a plant.
// Ideally the JSON and BSON tags should be defined in an specific entity
// for the API and storage components, respectively, but this being a small
// service they are defined here for simplicity.
type Plant struct {
	ID            string    `json:"id" bson:"_id"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	EditedAt      time.Time `json:"edited_at" bson:"edited_at"`
	Name          string    `json:"name" bson:"name"`
	OtherNames    string    `json:"other_names" bson:"other_names"`
	Description   string    `json:"description" bson:"description"`
	PlantSeason   string    `json:"plant_season" bson:"plant_season"`
	HarvestSeason string    `json:"harvest_season" bson:"harvest_season"`
	PruneSeason   string    `json:"prune_season" bson:"prune_season"`
	Tips          string    `json:"tips" bson:"tips"`
}

// AddService defines a service to add a plant.
type AddService interface {
	Add(string, Plant) (interface{}, error)
}

// ListService defines a service to list plants.
type ListService interface {
	ListAll(string) ([]Plant, error)
	ListOne(string, string) (Plant, error)
}

// EditService defines a service to edit a plant.
type EditService interface {
	Edit(string, string, Plant) (int64, error)
}

// DeleteService defines a service to delete a plant.
type DeleteService interface {
	Delete(string, string) (int64, error)
}

// Repository defines storage operations.
type Repository interface {
	InsertOne(string, Plant) (interface{}, error)
	FindAll(string) ([]Plant, error)
	FindOne(string, string) (Plant, error)
	UpdateOne(string, string, Plant) (int64, error)
	DeleteOne(string, string) (int64, error)
}
