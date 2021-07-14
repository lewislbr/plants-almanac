package tokenstore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const collection = "tokens"

type repository struct {
	db *firestore.Client
}

func NewRepository(db *firestore.Client) *repository {
	return &repository{db}
}

func (r *repository) Add(tokenID string) error {
	_, err := r.db.Collection(collection).Doc(tokenID).Set(context.Background(), map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("error adding token: %w", err)
	}

	return nil
}

func (r *repository) CheckExists(tokenID string) (bool, error) {
	_, err := r.db.Collection(collection).Doc(tokenID).Get(context.Background())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return false, nil
		}

		return false, fmt.Errorf("error checking token: %w", err)
	}

	return true, nil
}
