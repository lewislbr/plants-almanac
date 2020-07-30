package list

import (
	"plants/pkg/storage/mongodb"

	"github.com/graphql-go/graphql"
)

// Plant resolver
func Plant(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["_id"].(string)
	plant := mongodb.FindOne(id)

	return plant, nil
}

// Plants resolver
func Plants(p graphql.ResolveParams) (interface{}, error) {
	plants := mongodb.FindAll()

	return plants, nil
}
