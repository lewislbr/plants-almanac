package storage

import (
	"context"

	u "users/src/user"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// MongoDB provides methods to store data in MongoDB.
type MongoDB struct{}

// InsertOne adds a new user.
func (m *MongoDB) InsertOne(new u.User) (interface{}, error) {
	result, err := collection.InsertOne(context.Background(), new)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result.InsertedID, nil
}

// FindOne retuns the queried user.
func (m *MongoDB) FindOne(email string) (u.User, error) {
	filter := bson.M{"email": email}

	var result u.User

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return u.User{}, errors.Wrap(err, "")
	}

	return result, nil
}
