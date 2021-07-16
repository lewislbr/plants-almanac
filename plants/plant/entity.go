package plant

import "time"

type Plant struct {
	ID            string    `json:"id" bson:"_id"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at"`
	EditedAt      time.Time `json:"edited_at,omitempty" bson:"edited_at"`
	Name          string    `json:"name,omitempty" bson:"name"`
	OtherNames    string    `json:"other_names,omitempty" bson:"other_names"`
	Description   string    `json:"description,omitempty" bson:"description"`
	PlantSeason   string    `json:"plant_season,omitempty" bson:"plant_season"`
	HarvestSeason string    `json:"harvest_season,omitempty" bson:"harvest_season"`
	PruneSeason   string    `json:"prune_season,omitempty" bson:"prune_season"`
	Tips          string    `json:"tips,omitempty" bson:"tips"`
}
