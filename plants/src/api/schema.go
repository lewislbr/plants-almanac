package api

import (
	"plants/src/add"
	"plants/src/delete"
	"plants/src/edit"
	"plants/src/list"
	"plants/src/storage"

	"github.com/graphql-go/graphql"
)

var schema graphql.Schema
var plantType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Plant",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"created_at": &graphql.Field{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"edited_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"other_names": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"plant_season": &graphql.Field{
				Type: graphql.String,
			},
			"harvest_season": &graphql.Field{
				Type: graphql.String,
			},
			"prune_season": &graphql.Field{
				Type: graphql.String,
			},
			"tips": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var a = add.NewService(&storage.MongoDB{})
var l = list.NewService(&storage.MongoDB{})
var e = edit.NewService(&storage.MongoDB{})
var d = delete.NewService(&storage.MongoDB{})

func init() {
	queries := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"plants": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.NewList(plantType)),
				Description: "Lists all plants, returning an array of objects with the existing plants, or an empty array if there are none.",
				Resolve:     listPlants(l),
			},
			"plant": &graphql.Field{
				Type:        plantType,
				Description: "Lists a plant by using its ID, returning an object with the plant, or null if it does not exist.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: listPlant(l),
			},
		},
	})

	mutations := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"add": &graphql.Field{
				Type:        graphql.ID,
				Description: "Adds a plant by using a name, and any other value for the available fields, returning the newly created plant ID",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"other_names": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"plant_season": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"harvest_season": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"prune_season": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"tips": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: addPlant(a),
			},
			"edit": &graphql.Field{
				Type:        graphql.Int,
				Description: "Edits a plant by using its ID, and adding any new values for the existing fields, returning an integer with the numbers of plants edited.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"other_names": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"plant_season": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"harvest_season": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"prune_season": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"tips": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: editPlant(e, l),
			},
			"delete": &graphql.Field{
				Type:        graphql.Int,
				Description: "Deletes a plant by using its ID, returning an integer with the numbers of plants deleted.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: deletePlant(d),
			},
		},
	})

	schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queries,
		Mutation: mutations,
	})
}
