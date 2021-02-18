package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	u "users/src/user"
)

var isDevelopment = os.Getenv("MODE") == "development"

type handler struct {
	cr u.CreateService
	an u.AuthenticateService
	az u.AuthorizeService
	gn u.GenerateService
}

// NewHandler initializes a handler with the necessary dependencies.
func NewHandler(cr u.CreateService, an u.AuthenticateService, az u.AuthorizeService, gn u.GenerateService) handler {
	return handler{cr, an, az, gn}
}

func (h handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var new u.User

	json.NewDecoder(r.Body).Decode(&new)

	err := h.cr.Create(new)
	if err != nil {
		if err == u.ErrMissingData {
			http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

			return
		}
		if err == u.ErrUserExists {
			http.Error(w, u.ErrUserExists.Error(), http.StatusConflict)

			return
		}

		w.WriteHeader(http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h handler) LogInUser(w http.ResponseWriter, r *http.Request) {
	var cred u.Credentials

	json.NewDecoder(r.Body).Decode(&cred)

	jwt, err := h.an.Authenticate(cred)
	if err != nil {
		if err == u.ErrMissingData {
			http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

			return
		}
		if err == u.ErrNotFound {
			http.Error(w, u.ErrNotFound.Error(), http.StatusNotFound)

			return
		}
		if err == u.ErrInvalidPassword {
			http.Error(w, u.ErrInvalidPassword.Error(), http.StatusBadRequest)

			return
		}

		w.WriteHeader(http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	if isDevelopment {
		w.Header().Add("Set-Cookie", "st="+jwt+"; HttpOnly; Max-Age=604800")
		w.Header().Add("Set-Cookie", "te=true; Max-Age=604800")
	} else {
		w.Header().Add("Set-Cookie", "st="+jwt+"; Domain=plantdex.app; HttpOnly; Max-Age=604800; SameSite=Strict; Secure")
		w.Header().Add("Set-Cookie", "te=true; Domain=plantdex.app; Max-Age=604800; SameSite=Strict; Secure")
	}
}

func (h handler) AuthorizeUser(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	jwt := strings.Split(authHeader, " ")[1]
	uid, err := h.az.Authorize(jwt)
	if err != nil {
		if err == u.ErrMissingData {
			http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

			return
		}
		if err == u.ErrInvalidToken {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}

		w.WriteHeader(http.StatusInternalServerError)

		log.Printf("%+v\n", err)
	}

	io.WriteString(w, uid)
}

func (h handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var jwt string

	for _, cookie := range r.Cookies() {
		if cookie.Name == "st" {
			jwt = cookie.Value
		}
	}

	uid, err := h.az.Authorize(jwt)
	if err != nil {
		if err == u.ErrMissingData {
			http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

			return
		}
		if err == u.ErrInvalidToken {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}

		w.WriteHeader(http.StatusInternalServerError)

		log.Printf("%+v\n", err)
	}
	jwt, err = h.gn.GenerateJWT(uid)
	if err != nil {
		if err == u.ErrMissingData {
			http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

			return
		}

		w.WriteHeader(http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	if isDevelopment {
		w.Header().Add("Set-Cookie", "st="+jwt+"; HttpOnly; Max-Age=604800")
		w.Header().Add("Set-Cookie", "te=true; Max-Age=604800")
	} else {
		w.Header().Add("Set-Cookie", "st="+jwt+"; Domain=plantdex.app; HttpOnly; Max-Age=604800; SameSite=Strict; Secure")
		w.Header().Add("Set-Cookie", "te=true; Domain=plantdex.app; Max-Age=604800; SameSite=Strict; Secure")
	}
}
