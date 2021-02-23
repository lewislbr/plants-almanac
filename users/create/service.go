package create

import (
	"time"

	u "users/user"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type createService struct {
	r u.Repository
}

// NewCreateService initializes a create service with the necessary dependencies.
func NewCreateService(r u.Repository) *createService {
	return &createService{r}
}

// Create creates a new user.
func (cs *createService) Create(new u.User) error {
	if new.Name == "" || new.Email == "" || new.Password == "" {
		return u.ErrMissingData
	}

	_, err := cs.r.FindOne(new.Email)
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

	_, err = cs.r.InsertOne(new)
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}
