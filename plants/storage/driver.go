package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase(uri, db string) (*mongo.Database, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Plants database ready ✅")

	return client.Database(db), nil
}

func DisconnectDatabase(ctx context.Context, db *mongo.Database) error {
	fmt.Println("Disconnecting database...")

	return db.Client().Disconnect(ctx)
}
