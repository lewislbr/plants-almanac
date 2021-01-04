package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	u "users/src/user"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
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

// MongoDB provides methods to store data in MongoDB
type MongoDB struct{}

// InsertOne adds a new user
func (s *MongoDB) InsertOne(newUser u.User) (interface{}, error) {
	result, err := collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result.InsertedID, nil
}

// FindOne retuns the queried user
func (s *MongoDB) FindOne(email string) (u.User, error) {
	filter := bson.M{"email": email}

	var result u.User

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return u.User{}, errors.Wrap(err, "")
	}

	return result, nil
}
