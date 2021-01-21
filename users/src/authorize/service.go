package authorize

import (
	"os"

	u "users/src/user"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// Authorize checks if a user is authorized and returns its ID.
func Authorize(jwt string) (string, error) {
	if jwt == "" {
		return "", u.ErrMissingData
	}

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
