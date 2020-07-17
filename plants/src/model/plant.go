package model

// Plant entity
type Plant struct {
	ID            string  `json:"_id" bson:"_id"`
	Name          string  `json:"name" bson:"name"`
	OtherNames    *string `json:"otherNames" bson:"otherNames"`
	Description   *string `json:"description" bson:"description"`
	PlantSeason   *string `json:"plantSeason" bson:"plantSeason"`
	HarvestSeason *string `json:"harvestSeason" bson:"harvestSeason"`
	PruneSeason   *string `json:"pruneSeason" bson:"pruneSeason"`
	Tips          *string `json:"tips" bson:"tips"`
}
