package userstore

import (
	"context"
	"fmt"

	"lewislbr/plantdex/users/user"

	"cloud.google.com/go/firestore"
)

const collection = "users"

type repository struct {
	db *firestore.Client
}

func NewRepository(db *firestore.Client) *repository {
	return &repository{db}
}

func (r *repository) Insert(new user.User) error {
	_, err := r.db.Collection(collection).NewDoc().Create(context.Background(), new)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}

func (r *repository) Find(email string) (user.User, error) {
	res, err := r.db.Collection(collection).Where("email", "==", email).Limit(1).Documents(context.Background()).GetAll()
	if err != nil {
		return user.User{}, fmt.Errorf("error finding user: %w", err)
	}
	if res == nil {
		return user.User{}, nil
	}

	var data user.User

	err = res[0].DataTo(&data)
	if err != nil {
		return user.User{}, fmt.Errorf("error converting data: %w", err)
	}

	return data, nil
}

func (r *repository) CheckExists(email string) (bool, error) {
	res, err := r.db.Collection(collection).Where("email", "==", email).Limit(1).Documents(context.Background()).GetAll()
	if err != nil {
		return false, fmt.Errorf("error checking user: %w", err)
	}
	if res == nil {
		return false, nil
	}

	return true, nil
}

func (r *repository) GetInfo(userID string) (user.Info, error) {
	res, err := r.db.Collection(collection).Where("id", "==", userID).Select("name", "email", "created_at").Limit(1).Documents(context.Background()).GetAll()
	if err != nil {
		return user.Info{}, fmt.Errorf("error retrieving user info: %w", err)
	}
	if res == nil {
		return user.Info{}, nil
	}

	var info user.Info

	err = res[0].DataTo(&info)
	if err != nil {
		return user.Info{}, fmt.Errorf("error converting data: %w", err)
	}

	return info, nil
}
