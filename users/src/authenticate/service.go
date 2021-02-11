package authenticate

import (
	"os"
	"time"

	u "users/src/user"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type authenticateService struct {
	r u.Repository
}

// NewAuthenticateService initializes an authentication service with the necessary dependencies.
func NewAuthenticateService(r u.Repository) u.AuthenticateService {
	return authenticateService{r}
}

// Authenticate authenticates a user and issues a JWT.
func (s authenticateService) Authenticate(cred u.Credentials) (string, error) {
	if cred.Email == "" || cred.Password == "" {
		return "", u.ErrMissingData
	}

	existUser, err := s.r.FindOne(cred.Email)
	if err != nil {
		return "", u.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Hash), []byte(cred.Password))
	if err != nil {
		return "", u.ErrInvalidPassword
	}

	jwt, err := generateJWT(existUser.ID)
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	return jwt, nil
}

func generateJWT(uid string) (string, error) {
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
