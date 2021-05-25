package plant

import (
	"time"

	"github.com/google/uuid"
)

type (
	repository interface {
		Insert(string, Plant) (interface{}, error)
		FindAll(string) ([]Plant, error)
		FindOne(string, string) (Plant, error)
		Update(string, string, Plant) (int64, error)
		Delete(string, string) (int64, error)
	}

	service struct {
		repo repository
	}
)

func NewService(repo repository) *service {
	return &service{repo}
}

func (s *service) Add(userID string, new Plant) error {
	if new.Name == "" {
		return ErrMissingData
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()
	new.EditedAt = time.Now().UTC()

	_, err := s.repo.Insert(userID, new)

	return err
}

func (s *service) ListAll(userID string) ([]Plant, error) {
	result, err := s.repo.FindAll(userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) ListOne(userID, plantID string) (Plant, error) {
	if plantID == "" {
		return Plant{}, ErrMissingData
	}

	result, err := s.repo.FindOne(userID, plantID)
	if err != nil {
		return Plant{}, ErrNotFound
	}

	return result, nil
}

func (s *service) Edit(userID, plantID string, update Plant) error {
	if update.Name == "" {
		return ErrMissingData
	}

	exist, err := s.repo.FindOne(userID, plantID)
	if err != nil {
		return ErrNotFound
	}

	update.CreatedAt = exist.CreatedAt
	update.EditedAt = time.Now().UTC()

	result, err := s.repo.Update(userID, plantID, update)
	if err != nil {
		return err
	}
	if result == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *service) Delete(userID, plantID string) error {
	if plantID == "" {
		return ErrMissingData
	}

	result, err := s.repo.Delete(userID, plantID)
	if err != nil {
		return err
	}
	if result == 0 {
		return ErrNotFound
	}

	return nil
}
