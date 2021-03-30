package server

import (
	"encoding/json"
	"log"

	"plants/plant"

	"github.com/graphql-go/graphql"
)

type resolver struct {
	as plant.AddService
	ls plant.ListService
	es plant.EditService
	ds plant.DeleteService
}

func NewResolver(as plant.AddService, ls plant.ListService, es plant.EditService, ds plant.DeleteService) *resolver {
	return &resolver{as, ls, es, ds}
}

func (r *resolver) AddPlant(ps graphql.ResolveParams) (interface{}, error) {
	payload, err := json.Marshal(ps.Args)
	if err != nil {
		log.Println(err)
	}

	new := plant.Plant{}
	err = json.Unmarshal(payload, &new)
	if err != nil {
		log.Println(err)
	}

	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	result, err := r.as.Add(uid, new)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func (r *resolver) ListPlants(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	result, err := r.ls.ListAll(uid)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func (r *resolver) ListPlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	result, err := r.ls.ListOne(uid, id)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func (r *resolver) EditPlant(ps graphql.ResolveParams) (interface{}, error) {
	payload, err := json.Marshal(ps.Args)
	if err != nil {
		log.Println(err)
	}

	update := plant.Plant{}
	err = json.Unmarshal(payload, &update)
	if err != nil {
		log.Println(err)
	}

	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	result, err := r.es.Edit(uid, id, update)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}

func (r *resolver) DeletePlant(ps graphql.ResolveParams) (interface{}, error) {
	uid := ps.Info.RootValue.(map[string]interface{})["uid"].(string)
	id := ps.Args["id"].(string)
	result, err := r.ds.Delete(uid, id)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return result, nil
}
