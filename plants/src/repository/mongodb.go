package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"plants-go/src/model"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection a
func Connection() *mongo.Collection {
	godotenv.Load()

	MONGODB := os.Getenv("MONGODB_URI")
	databaseName := os.Getenv("DATABASE_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	clientOptions := options.Client().ApplyURI(MONGODB)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MongoDB database connected ✅\n")

	return client.Database(databaseName).Collection(collectionName)
}

var collection = Connection()

// FindAll returns all the plants
func FindAll() []*model.Plant {
	var plants []*model.Plant

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var plant *model.Plant
		err := cursor.Decode(&plant)
		if err != nil {
			log.Fatal(err)
		}

		plants = append(plants, plant)
	}

	return plants
}

// FindOne retuns the queried plant
func FindOne(filter bson.M) model.Plant {
	var plant model.Plant

	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&plant)

	return plant
}

// InsertOne adds a plant
func InsertOne(plant model.Plant) interface{} {
	result, err := collection.InsertOne(context.TODO(), plant)
	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID
}

// DeleteOne deletes a plant
func DeleteOne(filter bson.M) int64 {
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return deleteResult.DeletedCount
}
