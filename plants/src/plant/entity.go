package plant

import (
	"time"
)

// Plant defines the properties of a plant
type Plant struct {
	ID            ID        `json:"id" bson:"_id"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	EditedAt      time.Time `json:"edited_at" bson:"edited_at"`
	Name          string    `json:"name" bson:"name"`
	OtherNames    *string   `json:"other_names" bson:"other_names"`
	Description   *string   `json:"description" bson:"description"`
	PlantSeason   *string   `json:"plant_season" bson:"plant_season"`
	HarvestSeason *string   `json:"harvest_season" bson:"harvest_season"`
	PruneSeason   *string   `json:"prune_season" bson:"prune_season"`
	Tips          *string   `json:"tips" bson:"tips"`
}

// ID defines the id of a plant
type ID string
