package server

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"lewislbr/plantdex/users/token"
	"lewislbr/plantdex/users/user"
)

type (
	userService interface {
		Create(user.New) error
		Authenticate(cred user.Credentials) (string, error)
		Info(string) (user.Info, error)
	}
	tokenService interface {
		Generate(string) (string, error)
		Validate(string) (string, error)
		Revoke(string) error
	}

	handler struct {
		userSvc  userService
		tokenSvc tokenService
		domain   string
	}
)

func NewHandler(userSvc userService, tokenSvc tokenService, domain string) *handler {
	return &handler{userSvc, tokenSvc, domain}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var new user.New

	err := json.NewDecoder(r.Body).Decode(&new)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)

		log.Printf("Error decoding create request: %v\n", err)

		return
	}

	err = h.userSvc.Create(new)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrMissingData):
			http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, user.ErrUserExists):
			http.Error(w, user.ErrUserExists.Error(), http.StatusConflict)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error creating user: %v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var cred user.Credentials

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)

		log.Printf("Error decoding log in request: %v\n", err)

		return
	}

	userID, err := h.userSvc.Authenticate(cred)
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
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error logging in user: %v\n", err)

			return
		}
	}

	tkn, err := h.tokenSvc.Generate(userID)
	if err != nil {
		switch {
		case errors.Is(err, token.ErrMissingData):
			http.Error(w, token.ErrMissingData.Error(), http.StatusBadRequest)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error logging in user: %v\n", err)

			return
		}
	}

	w.Header().Add("Set-Cookie", "st="+tkn+"; Domain="+h.domain+"; HttpOnly; Max-Age=604800; Path=/; SameSite=Strict; Secure")
	w.Header().Add("Set-Cookie", "te=1; Domain="+h.domain+"; Max-Age=604800; Path=/; SameSite=Strict; Secure")

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) Authorize(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)

		log.Printf("Error reading authorization request: %v\n", err)

		return
	}

	tkn := string(body)
	userID, err := h.tokenSvc.Validate(tkn)
	if err != nil {
		switch {
		case errors.Is(err, token.ErrMissingData):
			http.Error(w, token.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, token.ErrInvalidToken):
			http.Error(w, token.ErrInvalidToken.Error(), http.StatusUnauthorized)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error authorizing user: %v\n", err)

			return
		}
	}

	_, err = io.WriteString(w, userID)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)

		log.Printf("Error authorizing user: %v\n", err)

		return
	}
}

func (h *handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var tkn string

	for _, cookie := range r.Cookies() {
		if cookie.Name == "st" {
			tkn = cookie.Value
		}
	}

	userID, err := h.tokenSvc.Validate(tkn)
	if err != nil {
		switch {
		case errors.Is(err, token.ErrMissingData):
			http.Error(w, token.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, token.ErrInvalidToken):
			http.Error(w, token.ErrInvalidToken.Error(), http.StatusUnauthorized)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error refreshing token: %v\n", err)

			return
		}
	}

	err = h.tokenSvc.Revoke(tkn)
	if err != nil {
		switch {
		case errors.Is(err, token.ErrMissingData):
			http.Error(w, token.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, token.ErrInvalidToken):
			http.Error(w, token.ErrInvalidToken.Error(), http.StatusUnauthorized)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error refreshing token: %v\n", err)

			return
		}
	}

	tkn, err = h.tokenSvc.Generate(userID)
	if err != nil {
		switch {
		case errors.Is(err, token.ErrMissingData):
			http.Error(w, token.ErrMissingData.Error(), http.StatusBadRequest)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error refreshing token: %v\n", err)

			return
		}
	}

	w.Header().Add("Set-Cookie", "st="+tkn+"; Domain="+h.domain+"; HttpOnly; Max-Age=604800; Path=/; SameSite=Strict; Secure")
	w.Header().Add("Set-Cookie", "te=1; Domain="+h.domain+"; Max-Age=604800; Path=/; SameSite=Strict; Secure")

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) LogOut(w http.ResponseWriter, r *http.Request) {
	var tkn string

	for _, cookie := range r.Cookies() {
		if cookie.Name == "st" {
			tkn = cookie.Value
		}
	}

	err := h.tokenSvc.Revoke(tkn)
	if err != nil {
		switch {
		case errors.Is(err, token.ErrMissingData):
			http.Error(w, token.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, token.ErrInvalidToken):
			http.Error(w, token.ErrInvalidToken.Error(), http.StatusUnauthorized)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error logging out user: %v\n", err)

			return
		}
	}

	w.Header().Add("Set-Cookie", "st=''; Domain="+h.domain+"; HttpOnly; Max-Age=0; Path=/; SameSite=Strict; Secure")
	w.Header().Add("Set-Cookie", "te=false; Domain="+h.domain+"; Max-Age=0; Path=/; SameSite=Strict; Secure")

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) Info(w http.ResponseWriter, r *http.Request) {
	var tkn string

	for _, cookie := range r.Cookies() {
		if cookie.Name == "st" {
			tkn = cookie.Value
		}
	}

	userID, err := h.tokenSvc.Validate(tkn)
	if err != nil {
		switch {
		case errors.Is(err, token.ErrMissingData):
			http.Error(w, token.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, token.ErrInvalidToken):
			http.Error(w, token.ErrInvalidToken.Error(), http.StatusUnauthorized)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error getting user info: %v\n", err)

			return
		}
	}

	data, err := h.userSvc.Info(userID)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrMissingData):
			http.Error(w, user.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, user.ErrNotFound):
			http.Error(w, user.ErrNotFound.Error(), http.StatusNotFound)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error getting user info: %v\n", err)

			return
		}
	}

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)

		log.Printf("Error encoding user info response: %v\n", err)

		return
	}
}
