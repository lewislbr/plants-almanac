package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	database *pgxpool.Pool
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(uri string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	dbpool, err := pgxpool.Connect(ctx, uri)
	if err != nil {
		return nil, err
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("Users database ready ✅")

	s.database = dbpool

	return s.database, nil
}

func (s *Storage) Disconnect() {
	fmt.Println("Disconnecting database...")

	s.database.Close()
}
