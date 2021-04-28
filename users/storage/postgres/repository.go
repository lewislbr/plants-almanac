package postgres

import (
	"context"

	"users/user"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{db}
}

func (r *repository) InsertOne(new user.User) error {
	_, err := r.db.Exec(context.Background(), `
		INSERT INTO user_(id, name, email, hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6);
		`,
		new.ID, new.Name, new.Email, new.Hash, new.CreatedAt, new.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindOne(email string) (user.User, error) {
	var result user.User

	row := r.db.QueryRow(context.Background(), `SELECT id, hash FROM user_ WHERE email = $1 LIMIT 1;`, email)
	err := row.Scan(&result.ID, &result.Hash)
	if err != nil {
		return user.User{}, err
	}

	return result, nil
}

func (r *repository) CheckExists(email string) (bool, error) {
	var result bool

	row := r.db.QueryRow(context.Background(), `SELECT EXISTS (SELECT 1 FROM user_ WHERE email = $1 LIMIT 1);`, email)
	err := row.Scan(&result)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (r *repository) GetUserInfo(id string) (user.Info, error) {
	var result user.Info

	row := r.db.QueryRow(context.Background(), `SELECT name, email, created_at FROM user_ WHERE id = $1 LIMIT 1;`, id)
	err := row.Scan(&result.Name, &result.Email, &result.CreatedAt)
	if err != nil {
		return user.Info{}, err
	}

	return result, nil
}
