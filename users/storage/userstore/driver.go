package userstore

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

type Driver struct {
	client *firestore.Client
}

func New() *Driver {
	return &Driver{}
}

func (d *Driver) Connect(projectID string) (*firestore.Client, error) {
	client, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		return nil, fmt.Errorf("error connecting user driver: %w", err)
	}

	log.Println("User database ready ✅")

	d.client = client

	return d.client, nil
}
