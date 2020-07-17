package schema

import (
	"plants/src/graphql/resolvers"

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
		"getPlant": &graphql.Field{
			Type:        plantType,
			Description: "Returns a plant",
			Args: graphql.FieldConfigArgument{
				"_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: resolvers.GetPlant,
		},
		"getPlants": &graphql.Field{
			Type:        graphql.NewList(plantType),
			Description: "Returns all plants",
			Resolve:     resolvers.GetPlants,
		},
	},
})

var mutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"addPlant": &graphql.Field{
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
			Resolve: resolvers.AddPlant,
		},
		"deletePlant": &graphql.Field{
			Type:        graphql.Int,
			Description: "Deletes a plant",
			Args: graphql.FieldConfigArgument{
				"_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: resolvers.DeletePlant,
		},
	},
})

// Schema data
var Schema, nil = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queries,
	Mutation: mutations,
})
