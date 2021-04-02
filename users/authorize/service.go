package authorize

import (
	"users/user"

	jwtgo "github.com/dgrijalva/jwt-go"
)

type service struct {
	secret string
}

func NewService(secret string) *service {
	return &service{
		secret: secret,
	}
}

func (s *service) Authorize(jwt string) (string, error) {
	if jwt == "" {
		return "", user.ErrMissingData
	}

	token, err := jwtgo.Parse(jwt, func(token *jwtgo.Token) (interface{}, error) {
		return []byte(s.secret), nil
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
