package delete

import (
	"plants/pkg/storage/mongodb"

	"github.com/graphql-go/graphql"
)

// Plant resolver
func Plant(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["_id"].(string)
	result := mongodb.DeleteOne(id)

	return result, nil
}
