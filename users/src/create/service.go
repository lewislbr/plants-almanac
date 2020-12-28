package create

import (
	"log"
	"time"

	u "users/src/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Service provides user creation operations
type Service interface {
	Create(u.User) error
}

// Repository provides access to the user storage
type Repository interface {
	FindOne(string) *u.User
	InsertOne(u.User) (interface{}, error)
}

type service struct {
	r Repository
}

// Create creates a new user
func (s *service) Create(newUser u.User) error {
	existUser := s.r.FindOne(newUser.Email)
	if existUser != nil {
		return u.ErrUserExists
	}

	newUser.ID = u.ID(uuid.New().String())
	newUser.CreatedAt = time.Now().UTC()

	hash, err2 := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err2 != nil {
		log.Println(err2)

		return err2
	}

	newUser.Hash = string(hash)

	_, err3 := s.r.InsertOne(newUser)
	if err3 != nil {
		log.Println(err3)

		return err3
	}

	return nil
}

// NewService creates a create service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
