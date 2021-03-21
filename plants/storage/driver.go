package storage

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase() (*mongo.Database, error) {
	mongodbURI := os.Getenv("PLANTS_MONGODB_URI")
	databaseName := os.Getenv("PLANTS_DATABASE_NAME")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Plants database ready ✅")

	return client.Database(databaseName), nil
}

func DisconnectDatabase(ctx context.Context, db *mongo.Database) error {
	fmt.Println("Disconnecting database...")

	err := db.Client().Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}
