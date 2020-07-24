package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"plants/src/model"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDatabase() *mongo.Collection {
	godotenv.Load()

	mongodbURI := os.Getenv("MONGODB_URI")
	databaseName := os.Getenv("DATABASE_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	clientOptions := options.Client().ApplyURI(mongodbURI)

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

var collection = connectDatabase()

// FindAll returns all the items
func FindAll() []*model.Plant {
	var result []*model.Plant

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var item *model.Plant
		err := cursor.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, item)
	}

	return result
}

// FindOne retuns the queried item
func FindOne(id string) *model.Plant {
	var result *model.Plant

	filter := bson.M{"_id": id}
	item := collection.FindOne(context.TODO(), filter)

	item.Decode(&result)

	return result
}

// InsertOne adds an item
func InsertOne(item model.Plant) interface{} {
	result, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID
}

// EditOne modifies the queried item
func EditOne(id string, updated model.Plant) int64 {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":          updated.Name,
			"otherNames":    updated.OtherNames,
			"description":   updated.Description,
			"plantSeason":   updated.PlantSeason,
			"harvestSeason": updated.HarvestSeason,
			"pruneSeason":   updated.PruneSeason,
			"tips":          updated.Tips,
		},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return result.ModifiedCount
}

// DeleteOne deletes an item
func DeleteOne(id string) int64 {
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount
}
