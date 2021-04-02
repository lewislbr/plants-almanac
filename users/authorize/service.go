package authorize

import (
	"users/user"

	jwtgo "github.com/dgrijalva/jwt-go"
)

type AuthorizeService interface {
	Authorize(string) (string, error)
}

type authorizeService struct {
	secret string
}

func NewAuthorizeService(secret string) *authorizeService {
	return &authorizeService{
		secret: secret,
	}
}

func (zs *authorizeService) Authorize(jwt string) (string, error) {
	if jwt == "" {
		return "", user.ErrMissingData
	}

	token, err := jwtgo.Parse(jwt, func(token *jwtgo.Token) (interface{}, error) {
		return []byte(zs.secret), nil
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
