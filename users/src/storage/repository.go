package storage

import (
	"context"

	u "users/src/user"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// InsertOne adds a new user.
func InsertOne(new u.User) (interface{}, error) {
	result, err := collection.InsertOne(context.Background(), new)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result.InsertedID, nil
}

// FindOne retuns the queried user.
func FindOne(email string) (u.User, error) {
	filter := bson.M{"email": email}

	var result u.User

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return u.User{}, errors.Wrap(err, "")
	}

	return result, nil
}
