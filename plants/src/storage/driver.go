package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDatabase() *mongo.Database {
	isDevelopment := os.Getenv("MODE") == "development"

	var mongodbURI string
	var databaseName string
	if isDevelopment {
		mongodbURI = os.Getenv("PLANTS_DEVELOPMENT_MONGODB_URI")
		databaseName = os.Getenv("PLANTS_DEVELOPMENT_DATABASE_NAME")
	} else {
		mongodbURI = os.Getenv("PLANTS_PRODUCTION_MONGODB_URI")
		databaseName = os.Getenv("PLANTS_PRODUCTION_DATABASE_NAME")
	}

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

var db = connectDatabase()
