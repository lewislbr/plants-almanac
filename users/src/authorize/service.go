package authorize

import (
	"os"

	u "users/src/user"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// Service defines a service to authorize a user.
type Service interface {
	Authorize(string) (string, error)
}

type service struct{}

// NewService creates an authorization service with the necessary dependencies.
func NewService() Service {
	return &service{}
}

// Authorize checks if a user is authorized and returns its ID.
func (s *service) Authorize(jwt string) (string, error) {
	token, err := jwtgo.Parse(jwt, func(token *jwtgo.Token) (interface{}, error) {
		return []byte(os.Getenv("USERS_JWT_SECRET")), nil
	})
	if !token.Valid {
		return "", u.ErrInvalidToken
	}
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	userID := token.Claims.(jwtgo.MapClaims)["uid"]

	return userID.(string), nil
}
