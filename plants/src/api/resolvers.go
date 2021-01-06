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
	new := p.Plant{}

	new.Name = ps.Args["name"].(string)

	if result, ok := ps.Args["other_names"].(string); ok {
		new.OtherNames = &result
	}
	if result, ok := ps.Args["description"].(string); ok {
		new.Description = &result
	}
	if result, ok := ps.Args["plant_season"].(string); ok {
		new.PlantSeason = &result
	}
	if result, ok := ps.Args["harvest_season"].(string); ok {
		new.HarvestSeason = &result
	}
	if result, ok := ps.Args["prune_season"].(string); ok {
		new.PruneSeason = &result
	}
	if result, ok := ps.Args["tips"].(string); ok {
		new.Tips = &result
	}

	result, err := addService.Add(uid, new)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func listPlants(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	result, err := listService.ListAll(uid)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func listPlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	result, err := listService.ListOne(uid, p.ID(id))
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func editPlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	exist, err := listService.ListOne(uid, p.ID(id))
	if err != nil {
		log.Printf("%+v\n", err)
	}

	update := p.Plant{}

	update.CreatedAt = exist.CreatedAt

	if result, ok := ps.Args["name"].(string); ok {
		update.Name = result
	} else {
		update.Name = exist.Name
	}
	if result, ok := ps.Args["other_names"].(string); ok {
		update.OtherNames = &result
	} else {
		update.OtherNames = exist.OtherNames
	}
	if result, ok := ps.Args["description"].(string); ok {
		update.Description = &result
	} else {
		update.Description = exist.Description
	}
	if result, ok := ps.Args["plant_season"].(string); ok {
		update.PlantSeason = &result
	} else {
		update.PlantSeason = exist.PlantSeason
	}
	if result, ok := ps.Args["harvest_season"].(string); ok {
		update.HarvestSeason = &result
	} else {
		update.HarvestSeason = exist.HarvestSeason
	}
	if result, ok := ps.Args["prune_season"].(string); ok {
		update.PruneSeason = &result
	} else {
		update.PruneSeason = exist.PruneSeason
	}
	if result, ok := ps.Args["tips"].(string); ok {
		update.Tips = &result
	} else {
		update.Tips = exist.Tips
	}

	result, err := editService.Edit(uid, p.ID(id), update)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func deletePlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	result, err := deleteService.Delete(uid, p.ID(id))
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}
