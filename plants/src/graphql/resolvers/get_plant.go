package resolvers

import (
	"plants-go/src/repository"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

// GetPlant resolver
func GetPlant(p graphql.ResolveParams) (interface{}, error) {
	return repository.FindOne(bson.M{"name": p.Args["name"]}), nil
}
