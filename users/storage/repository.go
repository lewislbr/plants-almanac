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

func NewRepository(db *mongo.Collection) *repository {
	return &repository{db}
}

func (r *repository) InsertOne(new u.User) (interface{}, error) {
	result, err := r.db.InsertOne(context.Background(), new)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (r *repository) FindOne(email string) (u.User, error) {
	filter := bson.M{"email": email}

	var result u.User

	err := r.db.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return u.User{}, err
	}

	return result, nil
}
