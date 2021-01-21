package create

import (
	"time"

	"users/src/storage"
	u "users/src/user"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Create creates a new user.
func Create(new u.User) error {
	if new.Name == "" || new.Email == "" || new.Password == "" {
		return u.ErrMissingData
	}

	_, err := storage.FindOne(new.Email)
	if err == nil {
		return u.ErrUserExists
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()

	hash, err := bcrypt.GenerateFromPassword([]byte(new.Password), 10)
	if err != nil {
		return errors.Wrap(err, "")
	}

	new.Hash = string(hash)

	_, err = storage.InsertOne(new)
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}
