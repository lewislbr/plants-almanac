package plant

import (
	"errors"
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

// ErrMissingData defines an error to use when there are missing required fields.
var ErrMissingData = errors.New("missing data")

// ErrNotFound defines an error to use when a plant is not found.
var ErrNotFound = errors.New("plant not found")
