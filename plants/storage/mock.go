package storage

import (
	"errors"

	"plants/plant"
)

type MockRepo struct {
	Plants []plant.Plant
}

func (m *MockRepo) InsertOne(uid string, new plant.Plant) (interface{}, error) {
	return new, nil
}

func (m *MockRepo) FindAll(uid string) ([]plant.Plant, error) {
	return []plant.Plant{}, nil
}

func (m *MockRepo) FindOne(uid, id string) (plant.Plant, error) {
	for _, p := range m.Plants {
		if id == p.ID {
			return p, nil
		}
	}

	return plant.Plant{}, errors.New("")
}

func (m *MockRepo) UpdateOne(uid, id string, update plant.Plant) (int64, error) {
	for _, p := range m.Plants {
		if id != p.ID {
			return 0, nil
		}
	}

	return 1, nil
}

func (m *MockRepo) DeleteOne(uid, id string) (int64, error) {
	for _, p := range m.Plants {
		if id != p.ID {
			return 0, nil
		}
	}

	return 1, nil
}
