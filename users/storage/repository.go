package storage

import (
	"context"

	u "users/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Collection
}

// NewRepository initializes a storage with the necessary dependencies.
func NewRepository(db *mongo.Collection) *repository {
	return &repository{db}
}

// InsertOne adds a new user.
func (r *repository) InsertOne(new u.User) (interface{}, error) {
	result, err := r.db.InsertOne(context.Background(), new)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

// FindOne retuns the queried user.
func (r *repository) FindOne(email string) (u.User, error) {
	filter := bson.M{"email": email}

	var result u.User

	err := r.db.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return u.User{}, err
	}

	return result, nil
}
