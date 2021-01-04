package api

import (
	"log"
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

	result, err := addService.AddPlant(uid, plant)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func listPlants(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	result, err := listService.ListPlants(uid)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func listPlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	result, err := listService.ListPlant(uid, p.ID(id))
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func editPlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	existing, err := listService.ListPlant(uid, p.ID(id))
	if err != nil {
		log.Printf("%+v\n", err)
	}

	updated := p.Plant{}

	updated.CreatedAt = existing.CreatedAt

	if result, ok := ps.Args["name"].(string); ok {
		updated.Name = result
	} else {
		updated.Name = existing.Name
	}
	if result, ok := ps.Args["other_names"].(string); ok {
		updated.OtherNames = &result
	} else {
		updated.OtherNames = existing.OtherNames
	}
	if result, ok := ps.Args["description"].(string); ok {
		updated.Description = &result
	} else {
		updated.Description = existing.Description
	}
	if result, ok := ps.Args["plant_season"].(string); ok {
		updated.PlantSeason = &result
	} else {
		updated.PlantSeason = existing.PlantSeason
	}
	if result, ok := ps.Args["harvest_season"].(string); ok {
		updated.HarvestSeason = &result
	} else {
		updated.HarvestSeason = existing.HarvestSeason
	}
	if result, ok := ps.Args["prune_season"].(string); ok {
		updated.PruneSeason = &result
	} else {
		updated.PruneSeason = existing.PruneSeason
	}
	if result, ok := ps.Args["tips"].(string); ok {
		updated.Tips = &result
	} else {
		updated.Tips = existing.Tips
	}

	result, err := editService.EditPlant(uid, p.ID(id), updated)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func deletePlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	result, err := deleteService.DeletePlant(uid, p.ID(id))
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}
