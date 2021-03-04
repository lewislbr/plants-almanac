package storage

import (
	"errors"
	p "plants/plant"
)

type MockRepo struct {
	Plants []p.Plant
}

func (m *MockRepo) InsertOne(uid string, new p.Plant) (interface{}, error) {
	return new, nil
}

func (m *MockRepo) FindAll(uid string) ([]p.Plant, error) {
	return []p.Plant{}, nil
}

func (m *MockRepo) FindOne(uid string, id string) (p.Plant, error) {
	for _, p := range m.Plants {
		if id == p.ID {
			return p, nil
		}
	}

	return p.Plant{}, errors.New("")
}

func (m *MockRepo) UpdateOne(uid string, id string, update p.Plant) (int64, error) {
	for _, p := range m.Plants {
		if id != p.ID {
			return 0, nil
		}
	}

	return 1, nil
}

func (m *MockRepo) DeleteOne(uid string, id string) (int64, error) {
	for _, p := range m.Plants {
		if id != p.ID {
			return 0, nil
		}
	}

	return 1, nil
}
