package resolvers

import (
	"plants/src/repository"

	"github.com/graphql-go/graphql"
)

// GetPlant resolver
func GetPlant(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["_id"].(string)
	plant := repository.FindOne(id)

	return plant, nil
}
