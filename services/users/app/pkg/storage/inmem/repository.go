package inmem

import (
	u "users/pkg/user"
)

func makeFakeDatabase() map[string]u.User {
	database := make(map[string]u.User)

	return database
}

var fakeDatabase = makeFakeDatabase()

// Storage provides methods to store data in memory
type Storage struct{}

// FindOne returns the queried user
func (s *Storage) FindOne(id u.ID) (*u.User, bool) {
	result, ok := fakeDatabase[string(id)]

	return &result, ok
}

// InsertOne adds a user
func (s *Storage) InsertOne(user u.User) {
	fakeDatabase[string(user.ID)] = user
}

// EditOne modifies the queried user
func (s *Storage) EditOne(id u.ID, user u.User) {
	fakeDatabase[string(id)] = user
}

// DeleteOne deletes a user
func (s *Storage) DeleteOne(id u.ID) {
	delete(fakeDatabase, string(id))
}
