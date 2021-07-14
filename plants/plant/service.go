package plant

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type (
	plantRepo interface {
		Insert(string, Plant) (interface{}, error)
		FindAll(string) ([]Plant, error)
		FindOne(string, string) (Plant, error)
		Update(string, string, Plant) (int64, error)
		Delete(string, string) (int64, error)
	}

	plantService struct {
		plantRepo plantRepo
	}
)

func NewService(plantRepo plantRepo) *plantService {
	return &plantService{plantRepo}
}

func (s *plantService) Add(userID string, new Plant) error {
	if new.Name == "" {
		return fmt.Errorf("error adding plant: %w", ErrMissingData)
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()
	new.EditedAt = time.Now().UTC()

	_, err := s.plantRepo.Insert(userID, new)
	if err != nil {
		return fmt.Errorf("error inserting plant: %w", err)
	}

	return nil
}

func (s *plantService) ListAll(userID string) ([]Plant, error) {
	result, err := s.plantRepo.FindAll(userID)
	if err != nil {
		return nil, fmt.Errorf("error finding plants: %w", err)
	}

	return result, nil
}

func (s *plantService) ListOne(userID, plantID string) (Plant, error) {
	if plantID == "" {
		return Plant{}, fmt.Errorf("error listing plant: %w", ErrMissingData)
	}

	result, err := s.plantRepo.FindOne(userID, plantID)
	if err != nil {
		return Plant{}, fmt.Errorf("error finding plant: %w", ErrNotFound)
	}

	return result, nil
}

func (s *plantService) Edit(userID, plantID string, update Plant) error {
	if update.Name == "" {
		return fmt.Errorf("error editing plant: %w", ErrMissingData)
	}

	exist, err := s.plantRepo.FindOne(userID, plantID)
	if err != nil {
		return fmt.Errorf("error finding plant: %w", ErrNotFound)
	}

	update.CreatedAt = exist.CreatedAt
	update.EditedAt = time.Now().UTC()

	result, err := s.plantRepo.Update(userID, plantID, update)
	if err != nil {
		return fmt.Errorf("error updating plant: %w", err)
	}
	if result == 0 {
		return fmt.Errorf("error updating plant: %w", ErrNotFound)
	}

	return nil
}

func (s *plantService) Delete(userID, plantID string) error {
	if plantID == "" {
		return fmt.Errorf("error deleting plant: %w", ErrMissingData)
	}

	result, err := s.plantRepo.Delete(userID, plantID)
	if err != nil {
		return fmt.Errorf("error deleting plant: %w", err)
	}
	if result == 0 {
		return fmt.Errorf("error deleting plant: %w", ErrNotFound)
	}

	return nil
}
