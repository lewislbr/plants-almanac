package inmem

import (
	"users/pkg/entity"
)

func makeFakeDatabase() map[string]entity.User {
	database := make(map[string]entity.User)

	return database
}

var fakeDatabase = makeFakeDatabase()

// Storage keeps data in memory
type Storage struct{}

// FindOne returns the queried user
func (s *Storage) FindOne(id string) (*entity.User, bool) {
	result, ok := fakeDatabase[id]

	return &result, ok
}

// InsertOne adds an user
func (s *Storage) InsertOne(user entity.User) {
	fakeDatabase[user.ID] = user
}

// EditOne modifies the queried user
func (s *Storage) EditOne(id string, user entity.User) {
	fakeDatabase[id] = user
}

// DeleteOne deletes an user
func (s *Storage) DeleteOne(id string) {
	delete(fakeDatabase, id)
}
