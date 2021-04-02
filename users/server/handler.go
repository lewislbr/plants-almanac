package server

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"users/user"
)

var isDevelopment = os.Getenv("MODE") == "development"

type (
	Creater interface {
		Create(user.User) error
	}
	Authenticater interface {
		Authenticate(cred user.Credentials) (string, error)
	}
	Authorizer interface {
		Authorize(string) (string, error)
	}
	Generater interface {
		GenerateJWT(string) (string, error)
	}

	handler struct {
		cs Creater
		ns Authenticater
		zs Authorizer
		gs Generater
	}
)

func NewHandler(cs Creater, ns Authenticater, zs Authorizer, gs Generater) *handler {
	return &handler{cs, ns, zs, gs}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var new user.User

	json.NewDecoder(r.Body).Decode(&new)

	err := h.cs.Create(new)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrMissingData):
			http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, user.ErrUserExists):
			http.Error(w, user.ErrUserExists.Error(), http.StatusConflict)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var cred user.Credentials

	json.NewDecoder(r.Body).Decode(&cred)

	jwt, err := h.ns.Authenticate(cred)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrMissingData):
			http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, user.ErrNotFound):
			http.Error(w, user.ErrNotFound.Error(), http.StatusNotFound)

			return
		case errors.Is(err, user.ErrInvalidPassword):
			http.Error(w, user.ErrInvalidPassword.Error(), http.StatusBadRequest)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	if isDevelopment {
		w.Header().Add("Set-Cookie", "st="+jwt+"; HttpOnly; Max-Age=604800")
		w.Header().Add("Set-Cookie", "te=true; Max-Age=604800")

		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Header().Add("Set-Cookie", "st="+jwt+"; Domain=plantdex.app; HttpOnly; Max-Age=604800; SameSite=Strict; Secure")
		w.Header().Add("Set-Cookie", "te=true; Domain=plantdex.app; Max-Age=604800; SameSite=Strict; Secure")

		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *handler) Authorize(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	jwt := strings.Split(authHeader, " ")[1]
	uid, err := h.zs.Authorize(jwt)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrMissingData):
			http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, user.ErrInvalidToken):
			w.WriteHeader(http.StatusUnauthorized)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	io.WriteString(w, uid)
}

func (h *handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var jwt string

	for _, cookie := range r.Cookies() {
		if cookie.Name == "st" {
			jwt = cookie.Value
		}
	}

	uid, err := h.zs.Authorize(jwt)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrMissingData):
			http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, user.ErrInvalidToken):
			w.WriteHeader(http.StatusUnauthorized)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	jwt, err = h.gs.GenerateJWT(uid)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrMissingData):
			http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	if isDevelopment {
		w.Header().Add("Set-Cookie", "st="+jwt+"; HttpOnly; Max-Age=604800")
		w.Header().Add("Set-Cookie", "te=true; Max-Age=604800")

		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Header().Add("Set-Cookie", "st="+jwt+"; Domain=plantdex.app; HttpOnly; Max-Age=604800; SameSite=Strict; Secure")
		w.Header().Add("Set-Cookie", "te=true; Domain=plantdex.app; Max-Age=604800; SameSite=Strict; Secure")

		w.WriteHeader(http.StatusNoContent)
	}
}
