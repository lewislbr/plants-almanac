package resolvers

import (
	"plants/src/model"
	"plants/src/repository"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// AddPlant resolver
func AddPlant(p graphql.ResolveParams) (interface{}, error) {
	plant := model.Plant{}

	plant.ID = uuid.New().String()
	plant.Name = p.Args["name"].(string)

	if result, ok := p.Args["otherNames"].(string); ok {
		plant.OtherNames = result
	}
	if result, ok := p.Args["description"].(string); ok {
		plant.Description = result
	}
	if result, ok := p.Args["plantSeason"].(string); ok {
		plant.PlantSeason = result
	}
	if result, ok := p.Args["harvestSeason"].(string); ok {
		plant.HarvestSeason = result
	}
	if result, ok := p.Args["pruneSeason"].(string); ok {
		plant.PruneSeason = result
	}
	if result, ok := p.Args["tips"].(string); ok {
		plant.Tips = result
	}

	return repository.InsertOne(plant), nil
}
