package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	p "plants/src/plant"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
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

// MongoDB provides methods to store data in MongoDB
type MongoDB struct{}

// InsertOne adds a plant
func (s *MongoDB) InsertOne(uid string, plant p.Plant) (interface{}, error) {
	result, err := db.Collection(uid).InsertOne(context.Background(), plant)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result.InsertedID, nil
}

// FindAll returns all the plants
func (s *MongoDB) FindAll(uid string) ([]p.Plant, error) {
	cursor, err := db.Collection(uid).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	var results []p.Plant

	err = cursor.All(context.Background(), &results)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return results, nil
}

// FindOne retuns the queried plant
func (s *MongoDB) FindOne(uid string, id p.ID) (p.Plant, error) {
	filter := bson.M{"_id": id}

	var result p.Plant

	err := db.Collection(uid).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return p.Plant{}, errors.Wrap(err, "")
	}

	return result, nil
}

// UpdateOne modifies the queried plant
func (s *MongoDB) UpdateOne(uid string, id p.ID, plant p.Plant) (int64, error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"created_at":     plant.CreatedAt,
			"edited_at":      plant.EditedAt,
			"name":           plant.Name,
			"other_names":    plant.OtherNames,
			"description":    plant.Description,
			"plant_season":   plant.PlantSeason,
			"harvest_season": plant.HarvestSeason,
			"prune_season":   plant.PruneSeason,
			"tips":           plant.Tips,
		},
	}
	result, err := db.Collection(uid).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result.ModifiedCount, nil
}

// DeleteOne deletes a plant
func (s *MongoDB) DeleteOne(uid string, id p.ID) (int64, error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(uid).DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result.DeletedCount, nil
}
