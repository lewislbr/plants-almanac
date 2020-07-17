package resolvers

import (
	"plants/src/repository"

	"github.com/graphql-go/graphql"
)

// GetPlants resolver
func GetPlants(p graphql.ResolveParams) (interface{}, error) {
	plants := repository.FindAll()

	return plants, nil
}
