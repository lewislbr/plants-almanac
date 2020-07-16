package resolvers

import (
	"plants-go/src/model"
	"plants-go/src/repository"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// AddPlant resolver
func AddPlant(p graphql.ResolveParams) (interface{}, error) {
	id := uuid.New().String()
	name := p.Args["name"].(string)
	otherNames, otherNamesOk := p.Args["otherNames"].(*string)
	if otherNamesOk {
		p.Args["otherNames"] = otherNames
	}
	description, descriptionOk := p.Args["description"].(*string)
	if descriptionOk {
		p.Args["description"] = description
	}
	plantSeason, plantSeasonOk := p.Args["plantSeason"].(*string)
	if plantSeasonOk {
		p.Args["plantSeason"] = plantSeason
	}
	harvestSeason, harvestSeasonOk := p.Args["harvestSeason"].(*string)
	if harvestSeasonOk {
		p.Args["harvestSeason"] = harvestSeason
	}
	pruneSeason, pruneSeasonOk := p.Args["pruneSeason"].(*string)
	if pruneSeasonOk {
		p.Args["pruneSeason"] = pruneSeason
	}
	tips, tipsOk := p.Args["tips"].(*string)
	if tipsOk {
		p.Args["tips"] = tips
	}

	plant := model.Plant{
		ID:            id,
		Name:          name,
		OtherNames:    otherNames,
		Description:   description,
		PlantSeason:   plantSeason,
		HarvestSeason: harvestSeason,
		PruneSeason:   pruneSeason,
		Tips:          tips,
	}

	return repository.InsertOne(plant), nil
}
