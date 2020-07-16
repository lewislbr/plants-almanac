package model

// Plant entity
type Plant struct {
	ID            string  `json:"_id" bson:"_id"`
	Name          string  `json:"name"`
	OtherNames    *string `json:"otherNames"`
	Description   *string `json:"description"`
	PlantSeason   *string `json:"plantSeason"`
	HarvestSeason *string `json:"harvestSeason"`
	PruneSeason   *string `json:"pruneSeason"`
	Tips          *string `json:"tips"`
}
