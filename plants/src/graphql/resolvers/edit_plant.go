package resolvers

import (
	"plants/src/model"
	"plants/src/repository"

	"github.com/graphql-go/graphql"
)

// EditPlant resolver
func EditPlant(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["_id"].(string)
	existingPlant := repository.FindOne(id)
	updated := model.Plant{}

	if result, ok := p.Args["name"].(string); ok {
		updated.Name = result
	} else {
		updated.Name = existingPlant.Name
	}
	if result, ok := p.Args["otherNames"].(string); ok {
		updated.OtherNames = &result
	} else {
		updated.OtherNames = existingPlant.OtherNames
	}
	if result, ok := p.Args["description"].(string); ok {
		updated.Description = &result
	} else {
		updated.Description = existingPlant.Description
	}
	if result, ok := p.Args["plantSeason"].(string); ok {
		updated.PlantSeason = &result
	} else {
		updated.PlantSeason = existingPlant.PlantSeason
	}
	if result, ok := p.Args["harvestSeason"].(string); ok {
		updated.HarvestSeason = &result
	} else {
		updated.HarvestSeason = existingPlant.HarvestSeason
	}
	if result, ok := p.Args["pruneSeason"].(string); ok {
		updated.PruneSeason = &result
	} else {
		updated.PruneSeason = existingPlant.PruneSeason
	}
	if result, ok := p.Args["tips"].(string); ok {
		updated.Tips = &result
	} else {
		updated.Tips = existingPlant.Tips
	}

	result := repository.EditOne(id, updated)

	return result, nil
}
