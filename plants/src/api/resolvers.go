package api

import (
	"plants/src/add"
	"plants/src/delete"
	"plants/src/edit"
	"plants/src/list"
	p "plants/src/plant"
	"plants/src/storage"

	"github.com/graphql-go/graphql"
)

var (
	addService    = add.NewService(&storage.MongoDB{})
	listService   = list.NewService(&storage.MongoDB{})
	editService   = edit.NewService(&storage.MongoDB{})
	deleteService = delete.NewService(&storage.MongoDB{})
)

func addPlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
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

	result := addService.AddPlant(uid, plant)

	return result, nil
}

func listPlants(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	plants := listService.ListPlants(uid)

	return plants, nil
}

func listPlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	plant := listService.ListPlant(uid, p.ID(id))

	return plant, nil
}

func editPlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	existingPlant := listService.ListPlant(uid, p.ID(id))
	updated := p.Plant{}

	updated.CreatedAt = existingPlant.CreatedAt

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

	result := editService.EditPlant(uid, p.ID(id), updated)

	return result, nil
}

func deletePlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	result := deleteService.DeletePlant(uid, p.ID(id))

	return result, nil
}
