package graphql

import (
	"plants/pkg/add"
	"plants/pkg/delete"
	"plants/pkg/edit"
	"plants/pkg/list"
	"plants/pkg/storage/mongodb"

	"github.com/graphql-go/graphql"
)

var l = list.NewService(&mongodb.Storage{})
var a = add.NewService(&mongodb.Storage{})
var e = edit.NewService(&mongodb.Storage{})
var d = delete.NewService(&mongodb.Storage{})
var plantType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Plant",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
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
var queries *graphql.Object
var mutations *graphql.Object
var schema graphql.Schema

func init() {
	queries = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"plants": &graphql.Field{
				Type:        graphql.NewList(plantType),
				Description: "Returns all plants",
				Resolve:     getPlants(l),
			},
			"plant": &graphql.Field{
				Type:        plantType,
				Description: "Returns a plant",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: getPlant(l),
			},
		},
	})

	mutations = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"add": &graphql.Field{
				Type:        graphql.ID,
				Description: "Adds a plant",
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
				Description: "Edits a plant",
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
				Description: "Deletes a plant",
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
