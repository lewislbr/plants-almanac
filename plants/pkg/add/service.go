package add

import (
	"plants/pkg/entity"
	"plants/pkg/storage/mongodb"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// Plant resolver
func Plant(p graphql.ResolveParams) (interface{}, error) {
	plant := entity.Plant{}

	plant.ID = uuid.New().String()
	plant.Name = p.Args["name"].(string)

	if result, ok := p.Args["otherNames"].(string); ok {
		plant.OtherNames = &result
	}
	if result, ok := p.Args["description"].(string); ok {
		plant.Description = &result
	}
	if result, ok := p.Args["plantSeason"].(string); ok {
		plant.PlantSeason = &result
	}
	if result, ok := p.Args["harvestSeason"].(string); ok {
		plant.HarvestSeason = &result
	}
	if result, ok := p.Args["pruneSeason"].(string); ok {
		plant.PruneSeason = &result
	}
	if result, ok := p.Args["tips"].(string); ok {
		plant.Tips = &result
	}

	result := mongodb.InsertOne(plant)

	return result, nil
}
