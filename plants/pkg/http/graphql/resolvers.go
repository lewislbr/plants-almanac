package graphql

import (
	"plants/pkg/add"
	"plants/pkg/delete"
	"plants/pkg/edit"
	"plants/pkg/entity"
	"plants/pkg/list"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

func getPlants(s list.Service) func(_ graphql.ResolveParams) (interface{}, error) {
	return func(_ graphql.ResolveParams) (interface{}, error) {
		plants := s.GetPlants()

		return plants, nil
	}
}

func getPlant(s list.Service) func(p graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["_id"].(string)
		plant := s.GetPlant(id)

		return plant, nil
	}
}

func addPlant(s add.Service) func(p graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
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

		result := s.AddPlant(plant)

		return result, nil
	}
}

func editPlant(s edit.Service, s2 list.Service) func(p graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["_id"].(string)
		existingPlant := s2.GetPlant(id)
		updated := entity.Plant{}

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

		result := s.EditPlant(id, updated)

		return result, nil
	}
}

func deletePlant(s delete.Service) func(p graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["_id"].(string)
		result := s.DeletePlant(id)

		return result, nil
	}
}
