package resolvers

import (
	"plants/src/repository"

	"github.com/graphql-go/graphql"
)

// DeletePlant resolver
func DeletePlant(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["_id"].(string)
	result := repository.DeleteOne(id)

	return result, nil
}
