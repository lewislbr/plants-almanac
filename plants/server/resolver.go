package server

import (
	"encoding/json"
	"log"

	p "plants/plant"

	"github.com/graphql-go/graphql"
)

type resolver struct {
	as p.AddService
	ls p.ListService
	es p.EditService
	ds p.DeleteService
}

// NewResolver initializes a handler with the necessary dependencies.
func NewResolver(as p.AddService, ls p.ListService, es p.EditService, ds p.DeleteService) *resolver {
	return &resolver{as, ls, es, ds}
}

func (r *resolver) AddPlant(ps graphql.ResolveParams) (interface{}, error) {
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

	update := p.Plant{}
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
