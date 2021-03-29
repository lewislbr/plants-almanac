package generate

import (
	"os"
	"time"

	"users/user"

	jwtgo "github.com/dgrijalva/jwt-go"
)

type GenerateService interface {
	GenerateJWT(string) (string, error)
}

type generateService struct{}

func NewGenerateService() *generateService {
	return &generateService{}
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
	secret := os.Getenv("USERS_JWT_SECRET")
	jwtString, err := jwt.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return jwtString, nil
}
