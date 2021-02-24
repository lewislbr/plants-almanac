package storage

import (
	u "users/user"
)

type MockRepo struct {
	Users []u.User
}

func (m *MockRepo) InsertOne(new u.User) (interface{}, error) {
	m.Users = append(m.Users, new)

	return new, nil
}

func (m *MockRepo) FindOne(email string) (u.User, error) {
	for _, u := range m.Users {
		if email == u.Email {
			return u, nil
		}
	}

	return u.User{}, u.ErrNotFound
}
