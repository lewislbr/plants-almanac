package graphql

import (
	"plants/pkg/add"
	"plants/pkg/delete"
	"plants/pkg/edit"
	"plants/pkg/list"

	"github.com/graphql-go/graphql"
)

var plantType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Plant",
		Fields: graphql.Fields{
			"_id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"otherNames": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"plantSeason": &graphql.Field{
				Type: graphql.String,
			},
			"harvestSeason": &graphql.Field{
				Type: graphql.String,
			},
			"pruneSeason": &graphql.Field{
				Type: graphql.String,
			},
			"tips": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"plant": &graphql.Field{
			Type:        plantType,
			Description: "Returns a plant",
			Args: graphql.FieldConfigArgument{
				"_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: list.Plant,
		},
		"plants": &graphql.Field{
			Type:        graphql.NewList(plantType),
			Description: "Returns all plants",
			Resolve:     list.Plants,
		},
	},
})

var mutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"add": &graphql.Field{
			Type:        graphql.ID,
			Description: "Adds a plant",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"otherNames": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"plantSeason": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"harvestSeason": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pruneSeason": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"tips": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: add.Plant,
		},
		"edit": &graphql.Field{
			Type:        graphql.Int,
			Description: "Edits a plant",
			Args: graphql.FieldConfigArgument{
				"_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"otherNames": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"plantSeason": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"harvestSeason": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pruneSeason": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"tips": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: edit.Plant,
		},
		"delete": &graphql.Field{
			Type:        graphql.Int,
			Description: "Deletes a plant",
			Args: graphql.FieldConfigArgument{
				"_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: delete.Plant,
		},
	},
})

var schema, nil = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queries,
	Mutation: mutations,
})
