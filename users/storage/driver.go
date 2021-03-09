package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDatabase establishes a connection with the database.
func ConnectDatabase() *mongo.Collection {
	mongodbURI := os.Getenv("USERS_MONGODB_URI")
	databaseName := os.Getenv("USERS_DATABASE_NAME")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users database ready ✅")

	collectionName := os.Getenv("USERS_COLLECTION_NAME")

	return client.Database(databaseName).Collection(collectionName)
}

// DisconnectDatabase closes the connection with the database.
func DisconnectDatabase(ctx context.Context, db *mongo.Collection) error {
	fmt.Println("Disconnecting database...")

	err := db.Database().Client().Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}
