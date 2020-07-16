package resolvers

import (
	"plants/src/repository"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

// DeletePlant resolver
func DeletePlant(p graphql.ResolveParams) (interface{}, error) {
	return repository.DeleteOne(bson.M{"_id": p.Args["_id"]}), nil
}
