package server

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	"users/user"
)

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
		GenerateToken(string) (string, error)
	}

	handler struct {
		cs     Creater
		ns     Authenticater
		zs     Authorizer
		gs     Generater
		domain string
	}
)

func NewHandler(cs Creater, ns Authenticater, zs Authorizer, gs Generater, domain string) *handler {
	return &handler{cs, ns, zs, gs, domain}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var new user.User

	err := json.NewDecoder(r.Body).Decode(&new)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	err = h.cs.Create(new)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrMissingData):
			http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, user.ErrUserExists):
			http.Error(w, user.ErrUserExists.Error(), http.StatusConflict)

			return
		default:
			http.Error(w, "something went wrong", http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var cred user.Credentials

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	token, err := h.ns.Authenticate(cred)
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
			http.Error(w, "something went wrong", http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.Header().Add("Set-Cookie", "st="+token+"; Domain="+h.domain+"; HttpOnly; Max-Age=604800; Path=/; SameSite=Strict; Secure")
	w.Header().Add("Set-Cookie", "te=true; Domain="+h.domain+"; Max-Age=604800; Path=/; SameSite=Strict; Secure")

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) Authorize(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

		return
	}

	token := strings.Split(authHeader, " ")[1]
	uid, err := h.zs.Authorize(token)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrMissingData):
			http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, user.ErrInvalidToken):
			http.Error(w, user.ErrInvalidToken.Error(), http.StatusUnauthorized)

			return
		default:
			http.Error(w, "something went wrong", http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	_, err = io.WriteString(w, uid)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}
}

func (h *handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var token string

	for _, cookie := range r.Cookies() {
		if cookie.Name == "st" {
			token = cookie.Value
		}
	}

	uid, err := h.zs.Authorize(token)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrMissingData):
			http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, user.ErrInvalidToken):
			http.Error(w, user.ErrInvalidToken.Error(), http.StatusUnauthorized)

			return
		default:
			http.Error(w, "something went wrong", http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	token, err = h.gs.GenerateToken(uid)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrMissingData):
			http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

			return
		default:
			http.Error(w, "something went wrong", http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.Header().Add("Set-Cookie", "st="+token+"; Domain="+h.domain+"; HttpOnly; Max-Age=604800; Path=/; SameSite=Strict; Secure")
	w.Header().Add("Set-Cookie", "te=true; Domain="+h.domain+"; Max-Age=604800; Path=/; SameSite=Strict; Secure")

	w.WriteHeader(http.StatusNoContent)
}
