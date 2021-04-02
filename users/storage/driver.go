package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase(uri, db, cl string) (*mongo.Collection, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Users database ready ✅")

	return client.Database(db).Collection(cl), nil
}

func DisconnectDatabase(ctx context.Context, db *mongo.Collection) error {
	fmt.Println("Disconnecting database...")

	return db.Database().Client().Disconnect(ctx)
}
