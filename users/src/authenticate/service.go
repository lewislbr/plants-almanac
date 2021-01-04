package authenticate

import (
	"os"

	u "users/src/user"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Service provides user authentication operations
type Service interface {
	Authenticate(cred u.Credentials) (string, error)
}

// Repository provides access to the user storage
type Repository interface {
	FindOne(string) (u.User, error)
}

type service struct {
	r Repository
}

// Authenticate authenticates a user and issues a JWT
func (s *service) Authenticate(cred u.Credentials) (string, error) {
	existUser, err := s.r.FindOne(cred.Email)
	if err != nil {
		return "", u.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Hash), []byte(cred.Password))
	if err != nil {
		return "", u.ErrInvalidPassword
	}

	jwt, err := generateJWT(existUser.ID)
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	return jwt, nil
}

func generateJWT(uid u.ID) (string, error) {
	jwt := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, jwtgo.MapClaims{
		"iss": "users",
		"uid": uid,
	})
	secret := os.Getenv("USERS_JWT_SECRET")
	jwtString, err := jwt.SignedString([]byte(secret))
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	return jwtString, nil
}

// NewService creates an authentication service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
