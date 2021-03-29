package authorize

import (
	"os"

	"users/user"

	jwtgo "github.com/dgrijalva/jwt-go"
)

type AuthorizeService interface {
	Authorize(string) (string, error)
}

type authorizeService struct{}

func NewAuthorizeService() *authorizeService {
	return &authorizeService{}
}

func (zs *authorizeService) Authorize(jwt string) (string, error) {
	if jwt == "" {
		return "", user.ErrMissingData
	}

	token, err := jwtgo.Parse(jwt, func(token *jwtgo.Token) (interface{}, error) {
		return []byte(os.Getenv("USERS_JWT_SECRET")), nil
	})
	if !token.Valid {
		return "", user.ErrInvalidToken
	}
	if err != nil {
		return "", err
	}

	userID := token.Claims.(jwtgo.MapClaims)["uid"]

	return userID.(string), nil
}
