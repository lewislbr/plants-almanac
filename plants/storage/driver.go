package storage

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase() (*mongo.Database, error) {
	ctx := context.Background()
	mongodbURI := os.Getenv("PLANTS_MONGODB_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Plants database ready ✅")

	databaseName := os.Getenv("PLANTS_DATABASE_NAME")

	return client.Database(databaseName), nil
}

func DisconnectDatabase(ctx context.Context, db *mongo.Database) error {
	fmt.Println("Disconnecting database...")

	return db.Client().Disconnect(ctx)
}
