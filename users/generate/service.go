package generate

import (
	"time"

	"users/user"

	jwtgo "github.com/dgrijalva/jwt-go"
)

type GenerateService interface {
	GenerateJWT(string) (string, error)
}

type generateService struct {
	secret string
}

func NewGenerateService(secret string) *generateService {
	return &generateService{
		secret: secret,
	}
}

func (gs *generateService) GenerateJWT(uid string) (string, error) {
	if uid == "" {
		return "", user.ErrMissingData
	}

	jwt := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, jwtgo.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iss": "users",
		"uid": uid,
	})
	jwtString, err := jwt.SignedString([]byte(gs.secret))
	if err != nil {
		return "", err
	}

	return jwtString, nil
}
