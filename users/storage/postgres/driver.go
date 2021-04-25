package postgres

import (
	"context"
	"fmt"

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
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("Users database ready ✅")

	d.pool = pool

	return d.pool, nil
}

func (d *Driver) Disconnect() {
	fmt.Println("Disconnecting database...")

	d.pool.Close()
}
