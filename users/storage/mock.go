package storage

import "users/user"

type MockRepo struct {
	Users []user.User
}

func (m *MockRepo) InsertOne(new user.User) (interface{}, error) {
	m.Users = append(m.Users, new)

	return new, nil
}

func (m *MockRepo) FindOne(email string) (user.User, error) {
	for _, u := range m.Users {
		if email == u.Email {
			return u, nil
		}
	}

	return user.User{}, user.ErrNotFound
}
