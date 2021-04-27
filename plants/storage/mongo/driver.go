package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Driver struct {
	database *mongo.Database
}

func New() *Driver {
	return &Driver{}
}

func (s *Driver) Connect(uri, db string) (*mongo.Database, error) {
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

	s.database = client.Database(db)

	return s.database, nil
}

func (s *Driver) Disconnect(ctx context.Context) error {
	if s.database == nil {
		return nil
	}

	fmt.Println("Disconnecting database...")

	return s.database.Client().Disconnect(ctx)
}
