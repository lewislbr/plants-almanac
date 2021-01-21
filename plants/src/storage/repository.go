package storage

import (
	"context"

	p "plants/src/plant"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// InsertOne adds a plant.
func InsertOne(uid string, new p.Plant) (interface{}, error) {
	result, err := db.Collection(uid).InsertOne(context.Background(), new)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result.InsertedID, nil
}

// FindAll returns all the plants.
func FindAll(uid string) ([]p.Plant, error) {
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

// FindOne retuns the queried plant.
func FindOne(uid string, id string) (p.Plant, error) {
	filter := bson.M{"_id": id}

	var result p.Plant

	err := db.Collection(uid).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return p.Plant{}, errors.Wrap(err, "")
	}

	return result, nil
}

// UpdateOne modifies the queried plant.
func UpdateOne(uid string, id string, update p.Plant) (int64, error) {
	filter := bson.M{"_id": id}
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
	result, err := db.Collection(uid).UpdateOne(context.Background(), filter, updated)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result.ModifiedCount, nil
}

// DeleteOne deletes a plant.
func DeleteOne(uid string, id string) (int64, error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(uid).DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result.DeletedCount, nil
}
