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
func ConnectDatabase() *mongo.Database {
	mongodbURI := os.Getenv("PLANTS_MONGODB_URI")
	databaseName := os.Getenv("PLANTS_DATABASE_NAME")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Plants database ready ✅")

	return client.Database(databaseName)
}
