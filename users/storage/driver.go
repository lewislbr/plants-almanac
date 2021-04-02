package storage

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase() (*mongo.Collection, error) {
	ctx := context.Background()
	mongodbURI := os.Getenv("USERS_MONGODB_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Users database ready ✅")

	databaseName := os.Getenv("USERS_DATABASE_NAME")
	collectionName := os.Getenv("USERS_COLLECTION_NAME")

	return client.Database(databaseName).Collection(collectionName), nil
}

func DisconnectDatabase(ctx context.Context, db *mongo.Collection) error {
	fmt.Println("Disconnecting database...")

	return db.Database().Client().Disconnect(ctx)
}
