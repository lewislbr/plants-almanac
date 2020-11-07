package graphql

import (
	"plants/pkg/add"
	"plants/pkg/delete"
	"plants/pkg/edit"
	"plants/pkg/list"
	p "plants/pkg/plant"

	"github.com/graphql-go/graphql"
)

func listPlants(ls list.Service) func(_ graphql.ResolveParams) (
	interface{},
	error,
) {
	return func(_ graphql.ResolveParams) (interface{}, error) {
		plants := ls.ListPlants()

		return plants, nil
	}
}

func listPlant(ls list.Service) func(ps graphql.ResolveParams) (
	interface{},
	error,
) {
	return func(ps graphql.ResolveParams) (interface{}, error) {
		id := ps.Args["id"].(string)
		plant := ls.ListPlant(p.ID(id))

		return plant, nil
	}
}

func addPlant(as add.Service) func(ps graphql.ResolveParams) (
	interface{},
	error,
) {
	return func(ps graphql.ResolveParams) (interface{}, error) {
		plant := p.Plant{}

		plant.Name = ps.Args["name"].(string)

		if result, ok := ps.Args["other_names"].(string); ok {
			plant.OtherNames = &result
		}
		if result, ok := ps.Args["description"].(string); ok {
			plant.Description = &result
		}
		if result, ok := ps.Args["plant_season"].(string); ok {
			plant.PlantSeason = &result
		}
		if result, ok := ps.Args["harvest_season"].(string); ok {
			plant.HarvestSeason = &result
		}
		if result, ok := ps.Args["prune_season"].(string); ok {
			plant.PruneSeason = &result
		}
		if result, ok := ps.Args["tips"].(string); ok {
			plant.Tips = &result
		}

		result := as.AddPlant(plant)

		return result, nil
	}
}

func editPlant(
	es edit.Service,
	ls list.Service,
) func(ps graphql.ResolveParams) (
	interface{},
	error,
) {
	return func(ps graphql.ResolveParams) (interface{}, error) {
		id := ps.Args["id"].(string)
		existingPlant := ls.ListPlant(p.ID(id))
		updated := p.Plant{}

		if result, ok := ps.Args["name"].(string); ok {
			updated.Name = result
		} else {
			updated.Name = existingPlant.Name
		}
		if result, ok := ps.Args["other_names"].(string); ok {
			updated.OtherNames = &result
		} else {
			updated.OtherNames = existingPlant.OtherNames
		}
		if result, ok := ps.Args["description"].(string); ok {
			updated.Description = &result
		} else {
			updated.Description = existingPlant.Description
		}
		if result, ok := ps.Args["plant_season"].(string); ok {
			updated.PlantSeason = &result
		} else {
			updated.PlantSeason = existingPlant.PlantSeason
		}
		if result, ok := ps.Args["harvest_season"].(string); ok {
			updated.HarvestSeason = &result
		} else {
			updated.HarvestSeason = existingPlant.HarvestSeason
		}
		if result, ok := ps.Args["prune_season"].(string); ok {
			updated.PruneSeason = &result
		} else {
			updated.PruneSeason = existingPlant.PruneSeason
		}
		if result, ok := ps.Args["tips"].(string); ok {
			updated.Tips = &result
		} else {
			updated.Tips = existingPlant.Tips
		}

		result := es.EditPlant(p.ID(id), updated)

		return result, nil
	}
}

func deletePlant(ds delete.Service) func(ps graphql.ResolveParams) (
	interface{},
	error,
) {
	return func(ps graphql.ResolveParams) (interface{}, error) {
		id := ps.Args["id"].(string)
		result := ds.DeletePlant(p.ID(id))

		return result, nil
	}
}
