package storage

import (
	"context"

	"users/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertOne(user.User) (interface{}, error)
	FindOne(string) (user.User, error)
}

type repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Collection) *repository {
	return &repository{db}
}

func (r *repository) InsertOne(new user.User) (interface{}, error) {
	result, err := r.db.InsertOne(context.Background(), new)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (r *repository) FindOne(email string) (user.User, error) {
	filter := bson.M{"email": email}

	var result user.User

	err := r.db.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return user.User{}, err
	}

	return result, nil
}
