package create

import (
	"time"

	"users/storage"
	"users/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateService interface {
	Create(user.User) error
}

type createService struct {
	r storage.Repository
}

func NewCreateService(r storage.Repository) *createService {
	return &createService{r}
}

func (cs *createService) Create(new user.User) error {
	if new.Name == "" || new.Email == "" || new.Password == "" {
		return user.ErrMissingData
	}

	_, err := cs.r.FindOne(new.Email)
	if err == nil {
		return user.ErrUserExists
	}

	new.ID = uuid.New().String()
	new.CreatedAt = time.Now().UTC()

	hash, err := bcrypt.GenerateFromPassword([]byte(new.Password), 10)
	if err != nil {
		return err
	}

	new.Hash = string(hash)

	_, err = cs.r.InsertOne(new)
	if err != nil {
		return err
	}

	return nil
}
