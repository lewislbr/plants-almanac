package authorize

import (
	"log"
	"os"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// Service provides user authorization operations
type Service interface {
	Authorize(string) string
}

type service struct{}

// Authorize checks if a user is authorized and returns its ID
func (s *service) Authorize(jwt string) string {
	token, err := jwtgo.Parse(jwt, func(token *jwtgo.Token) (interface{}, error) {
		return []byte(os.Getenv("USERS_JWT_SECRET")), nil
	})
	if err != nil {
		log.Println(err)

		return ""
	}
	if !token.Valid {
		return ""
	}

	userID := token.Claims.(jwtgo.MapClaims)["uid"]

	return userID.(string)
}

// NewService creates an authorization service with the necessary dependencies
func NewService() Service {
	return &service{}
}
