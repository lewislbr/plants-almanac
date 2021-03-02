package generate

import (
	"os"
	"time"

	u "users/user"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type generateService struct{}

// NewGenerateService initializes a generate service with the necessary dependencies.
func NewGenerateService() *generateService {
	return &generateService{}
}

func (gs *generateService) GenerateJWT(uid string) (string, error) {
	if uid == "" {
		return "", u.ErrMissingData
	}

	jwt := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, jwtgo.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
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