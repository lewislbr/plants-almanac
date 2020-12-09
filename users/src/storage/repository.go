package storage

import (
	u "users/src/user"
)

func makeFakeDatabase() map[string]u.User {
	database := make(map[string]u.User)

	return database
}

var fakeDatabase = makeFakeDatabase()

// Inmem provides methods to store data in memory
type Inmem struct{}

// FindOne returns the queried user
func (s *Inmem) FindOne(id u.ID) (*u.User, bool) {
	result, ok := fakeDatabase[string(id)]

	return &result, ok
}

// InsertOne adds a user
func (s *Inmem) InsertOne(user u.User) {
	fakeDatabase[string(user.ID)] = user
}

// EditOne modifies the queried user
func (s *Inmem) EditOne(id u.ID, user u.User) {
	fakeDatabase[string(id)] = user
}

// DeleteOne deletes a user
func (s *Inmem) DeleteOne(id u.ID) {
	delete(fakeDatabase, string(id))
}
