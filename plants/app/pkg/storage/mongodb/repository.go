package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"plants/pkg/entity"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var isDevelopment = os.Getenv("MODE") == "development"

func connectDatabase() *mongo.Collection {
	godotenv.Load()

	var mongodbURI string
	if isDevelopment {
		mongodbURI = os.Getenv("PLANTS_MONGODB_DEVELOPMENT_URI")
	} else {
		mongodbURI = os.Getenv("PLANTS_MONGODB_PRODUCTION_URI")
	}

	databaseName := os.Getenv("PLANTS_DATABASE_NAME")
	collectionName := os.Getenv("PLANTS_COLLECTION_NAME")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodbURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database ready ✅")

	return client.Database(databaseName).Collection(collectionName)
}

var collection = connectDatabase()

// Storage stores data in MongoDB Atlas
type Storage struct{}

// FindAll returns all the items
func (s *Storage) FindAll() []*entity.Plant {
	var result []*entity.Plant

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var item *entity.Plant
		err := cursor.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, item)
	}

	return result
}

// FindOne retuns the queried item
func (s *Storage) FindOne(id string) *entity.Plant {
	var result *entity.Plant

	filter := bson.M{"_id": id}
	item := collection.FindOne(context.TODO(), filter)

	item.Decode(&result)

	return result
}

// InsertOne adds an item
func (s *Storage) InsertOne(item entity.Plant) interface{} {
	result, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID
}

// EditOne modifies the queried item
func (s *Storage) EditOne(id string, updated entity.Plant) int64 {
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
func (s *Storage) DeleteOne(id string) int64 {
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount
}
