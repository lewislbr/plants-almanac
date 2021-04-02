package generate

import (
	"time"

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

func (s *service) GenerateJWT(uid string) (string, error) {
	if uid == "" {
		return "", user.ErrMissingData
	}

	jwt := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, jwtgo.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iss": "users",
		"uid": uid,
	})
	jwtString, err := jwt.SignedString([]byte(s.secret))
	if err != nil {
		return "", err
	}

	return jwtString, nil
}
