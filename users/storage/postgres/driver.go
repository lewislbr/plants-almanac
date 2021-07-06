package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Driver struct {
	pool *pgxpool.Pool
}

func New() *Driver {
	return &Driver{}
}

func (d *Driver) Connect(uri string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, uri)
	if err != nil {
		return nil, fmt.Errorf("error connecting Postgres driver: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("error pinging Postgres client: %w", err)
	}

	log.Println("Users database ready ✅")

	d.pool = pool

	return d.pool, nil
}

func (d *Driver) Disconnect() {
	log.Println("Disconnecting database...")

	d.pool.Close()
}
