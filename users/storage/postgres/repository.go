package postgres

import (
	"context"
	"fmt"

	"users/user"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{db}
}

func (r *repository) Insert(new user.User) error {
	_, err := r.db.Exec(context.Background(), `
		INSERT INTO user_(id, name, email, hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6);
		`,
		new.ID, new.Name, new.Email, new.Hash, new.CreatedAt, new.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}

func (r *repository) Find(email string) (user.User, error) {
	var result user.User

	row := r.db.QueryRow(context.Background(), `SELECT id, hash FROM user_ WHERE email = $1 LIMIT 1;`, email)
	err := row.Scan(&result.ID, &result.Hash)
	if err != nil {
		return user.User{}, fmt.Errorf("error finding user: %w", err)
	}

	return result, nil
}

func (r *repository) CheckExists(email string) (bool, error) {
	var result bool

	row := r.db.QueryRow(context.Background(), `SELECT EXISTS (SELECT 1 FROM user_ WHERE email = $1 LIMIT 1);`, email)
	err := row.Scan(&result)
	if err != nil {
		return false, fmt.Errorf("error checking user: %w", err)
	}

	return result, nil
}

func (r *repository) GetInfo(userID string) (user.Info, error) {
	var result user.Info

	row := r.db.QueryRow(context.Background(), `SELECT name, email, created_at FROM user_ WHERE id = $1 LIMIT 1;`, userID)
	err := row.Scan(&result.Name, &result.Email, &result.CreatedAt)
	if err != nil {
		return user.Info{}, fmt.Errorf("error retrieving user info: %w", err)
	}

	return result, nil
}
