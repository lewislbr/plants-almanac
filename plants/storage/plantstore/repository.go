package plantstore

import (
	"context"
	"fmt"

	"lewislbr/plantdex/plants/plant"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *repository {
	return &repository{db}
}

func (r *repository) Insert(userID string, new plant.Plant) (interface{}, error) {
	result, err := r.db.Collection(userID).InsertOne(context.Background(), new)
	if err != nil {
		return nil, fmt.Errorf("error inserting plant: %w", err)
	}

	return result.InsertedID, nil
}

func (r *repository) FindAll(userID string) ([]plant.Plant, error) {
	cursor, err := r.db.Collection(userID).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error finding plants: %w", err)
	}

	var results []plant.Plant

	err = cursor.All(context.Background(), &results)
	if err != nil {
		return nil, fmt.Errorf("error iterating plants: %w", err)
	}

	return results, nil
}

func (r *repository) FindOne(userID, plantID string) (plant.Plant, error) {
	filter := bson.M{"_id": plantID}

	var result plant.Plant

	err := r.db.Collection(userID).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return plant.Plant{}, fmt.Errorf("error finding plant: %w", plant.ErrNotFound)
		}

		return plant.Plant{}, fmt.Errorf("error finding plant: %w", err)
	}

	return result, nil
}

func (r *repository) Update(userID, plantID string, update plant.Plant) (int64, error) {
	filter := bson.M{"_id": plantID}
	updated := bson.M{
		"$set": bson.M{
			"created_at":     update.CreatedAt,
			"edited_at":      update.EditedAt,
			"name":           update.Name,
			"other_names":    update.OtherNames,
			"description":    update.Description,
			"plant_season":   update.PlantSeason,
			"harvest_season": update.HarvestSeason,
			"prune_season":   update.PruneSeason,
			"tips":           update.Tips,
		},
	}
	result, err := r.db.Collection(userID).UpdateOne(context.Background(), filter, updated)
	if err != nil {
		return 0, fmt.Errorf("error editing plant: %w", err)
	}

	return result.ModifiedCount, nil
}

func (r *repository) Delete(userID, plantID string) (int64, error) {
	filter := bson.M{"_id": plantID}
	result, err := r.db.Collection(userID).DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, fmt.Errorf("error deleting plant: %w", err)
	}

	return result.DeletedCount, nil
}
