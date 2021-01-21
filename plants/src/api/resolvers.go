package api

import (
	"encoding/json"
	"log"
	"plants/src/add"
	"plants/src/delete"
	"plants/src/edit"
	"plants/src/list"
	p "plants/src/plant"

	"github.com/graphql-go/graphql"
)

func addPlant(ps graphql.ResolveParams) (interface{}, error) {
	payload, err := json.Marshal(ps.Args)
	if err != nil {
		log.Println(err)
	}

	new := p.Plant{}
	err = json.Unmarshal(payload, &new)
	if err != nil {
		log.Println(err)
	}

	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	result, err := add.Add(uid, new)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func listPlants(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	result, err := list.ListAll(uid)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func listPlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	result, err := list.ListOne(uid, id)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func editPlant(ps graphql.ResolveParams) (interface{}, error) {
	payload, err := json.Marshal(ps.Args)
	if err != nil {
		log.Println(err)
	}

	update := p.Plant{}
	err = json.Unmarshal(payload, &update)
	if err != nil {
		log.Println(err)
	}

	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	result, err := edit.Edit(uid, id, update)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func deletePlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	result, err := delete.Delete(uid, id)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}
