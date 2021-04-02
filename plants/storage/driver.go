package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

func Connect(uri, db string) (*mongo.Database, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Plants database ready ✅")

	database = client.Database(db)

	return database, nil
}

func Disconnect(ctx context.Context) error {
	if database == nil {
		return nil
	}

	fmt.Println("Disconnecting database...")

	return database.Client().Disconnect(ctx)
}
