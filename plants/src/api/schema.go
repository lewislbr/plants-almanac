package api

import (
	"github.com/graphql-go/graphql"
)

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

// NewSchema initializes the schema with the necessary dependencies.
func NewSchema(r resolver) (*graphql.Schema, error) {
	queries := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"plants": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.NewList(plantType)),
				Description: "Lists all plants, returning an array of objects with the existing plants, or an empty array if there are none.",
				Resolve:     r.ListPlants,
			},
			"plant": &graphql.Field{
				Type:        plantType,
				Description: "Lists a plant by using its ID, returning an object with the plant, or null if it does not exist.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: r.ListPlant,
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
				Resolve: r.AddPlant,
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
				Resolve: r.EditPlant,
			},
			"delete": &graphql.Field{
				Type:        graphql.Int,
				Description: "Deletes a plant by using its ID, returning an integer with the numbers of plants deleted.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: r.DeletePlant,
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queries,
		Mutation: mutations,
	})
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
