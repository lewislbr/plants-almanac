package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDatabase() *mongo.Collection {
	isDevelopment := os.Getenv("MODE") == "development"

	var mongodbURI string
	var databaseName string
	if isDevelopment {
		mongodbURI = os.Getenv("USERS_DEVELOPMENT_MONGODB_URI")
		databaseName = os.Getenv("USERS_DEVELOPMENT_DATABASE_NAME")
	} else {
		mongodbURI = os.Getenv("USERS_PRODUCTION_MONGODB_URI")
		databaseName = os.Getenv("USERS_PRODUCTION_DATABASE_NAME")
	}

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

var collection = connectDatabase()
