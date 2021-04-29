package server

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"users/user"
)

type (
	creater interface {
		Create(user.User) error
	}
	authenticater interface {
		Authenticate(cred user.Credentials) (string, error)
	}
	authorizer interface {
		Authorize(string) (string, error)
	}
	generater interface {
		GenerateToken(string) (string, error)
	}
	revoker interface {
		RevokeToken(string) error
	}
	infoer interface {
		UserInfo(string) (user.Info, error)
	}

	handler struct {
		createSvc       creater
		authenticateSvc authenticater
		authorizeSvc    authorizer
		generateSvc     generater
		revokeSvc       revoker
		infoSvc         infoer
		domain          string
	}
)

func NewHandler(
	createSvc creater,
	authenticateSvc authenticater,
	authorizeSvc authorizer,
	generateSvc generater,
	revokeSvc revoker,
	infoSvc infoer,
	domain string,
) *handler {
	return &handler{createSvc, authenticateSvc, authorizeSvc, generateSvc, revokeSvc, infoSvc, domain}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var new user.User

	err := json.NewDecoder(r.Body).Decode(&new)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	err = h.createSvc.Create(new)
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

	userID, err := h.authenticateSvc.Authenticate(cred)
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

	token, err := h.generateSvc.GenerateToken(userID)
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
	w.Header().Add("Set-Cookie", "te=1; Domain="+h.domain+"; Max-Age=604800; Path=/; SameSite=Strict; Secure")

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) Authorize(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	token := string(body)
	userID, err := h.authorizeSvc.Authorize(token)
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

	_, err = io.WriteString(w, userID)
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

	userID, err := h.authorizeSvc.Authorize(token)
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

	err = h.revokeSvc.RevokeToken(token)
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

	token, err = h.generateSvc.GenerateToken(userID)
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
	w.Header().Add("Set-Cookie", "te=1; Domain="+h.domain+"; Max-Age=604800; Path=/; SameSite=Strict; Secure")

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) LogOut(w http.ResponseWriter, r *http.Request) {
	var token string

	for _, cookie := range r.Cookies() {
		if cookie.Name == "st" {
			token = cookie.Value
		}
	}

	err := h.revokeSvc.RevokeToken(token)
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

	w.Header().Add("Set-Cookie", "st=''; Domain="+h.domain+"; HttpOnly; Max-Age=0; Path=/; SameSite=Strict; Secure")
	w.Header().Add("Set-Cookie", "te=false; Domain="+h.domain+"; Max-Age=0; Path=/; SameSite=Strict; Secure")

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) Info(w http.ResponseWriter, r *http.Request) {
	var token string

	for _, cookie := range r.Cookies() {
		if cookie.Name == "st" {
			token = cookie.Value
		}
	}

	userID, err := h.authorizeSvc.Authorize(token)
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

	data, err := h.infoSvc.UserInfo(userID)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrMissingData):
			http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, user.ErrNotFound):
			http.Error(w, user.ErrNotFound.Error(), http.StatusNotFound)

			return
		default:
			http.Error(w, "something went wrong", http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}
}
