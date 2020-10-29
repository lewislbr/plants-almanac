package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	p "plants/pkg/plant"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDatabase() *mongo.Collection {
	godotenv.Load()

	var isDevelopment = os.Getenv("MODE") == "development"
	var mongodbURI string
	if isDevelopment {
		mongodbURI = os.Getenv("PLANTS_MONGODB_DEVELOPMENT_URI")
	} else {
		mongodbURI = os.Getenv("PLANTS_MONGODB_PRODUCTION_URI")
	}

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(mongodbURI),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Plants database ready ✅")

	databaseName := os.Getenv("PLANTS_DATABASE_NAME")
	collectionName := os.Getenv("PLANTS_COLLECTION_NAME")

	return client.Database(databaseName).Collection(collectionName)
}

var collection = connectDatabase()

// Storage provides methods to store data in MongoDB
type Storage struct{}

// FindAll returns all the plants
func (s *Storage) FindAll() []*p.Plant {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var results []*p.Plant

	if err := cursor.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}

	return results
}

// FindOne retuns the queried plant
func (s *Storage) FindOne(id p.ID) *p.Plant {
	filter := bson.M{"_id": id}
	singleResult := collection.FindOne(context.Background(), filter)

	var result *p.Plant

	singleResult.Decode(&result)

	return result
}

// InsertOne adds a plant
func (s *Storage) InsertOne(plant p.Plant) interface{} {
	result, err := collection.InsertOne(context.Background(), plant)
	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID
}

// EditOne modifies the queried plant
func (s *Storage) EditOne(id p.ID, plant p.Plant) int64 {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":           plant.Name,
			"other_names":    plant.OtherNames,
			"description":    plant.Description,
			"plant_season":   plant.PlantSeason,
			"harvest_season": plant.HarvestSeason,
			"prune_season":   plant.PruneSeason,
			"tips":           plant.Tips,
		},
	}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return result.ModifiedCount
}

// DeleteOne deletes a plant
func (s *Storage) DeleteOne(id p.ID) int64 {
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount
}
