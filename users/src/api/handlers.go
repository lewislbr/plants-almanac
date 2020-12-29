package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"users/src/authenticate"
	"users/src/authorize"
	"users/src/create"
	"users/src/storage"
	u "users/src/user"
)

var createService = create.NewService(&storage.MongoDB{})
var authenticateService = authenticate.NewService(&storage.MongoDB{})
var authorizeService = authorize.NewService()

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser u.User

	json.NewDecoder(r.Body).Decode(&newUser)

	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" {
		http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

		return
	}

	err := createService.Create(newUser)
	if err != nil {
		if strings.Contains(err.Error(), u.ErrUserExists.Error()) {
			http.Error(w, u.ErrUserExists.Error(), http.StatusConflict)

			return
		}

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func logInUser(w http.ResponseWriter, r *http.Request) {
	var credentials u.Credentials

	json.NewDecoder(r.Body).Decode(&credentials)

	if credentials.Email == "" || credentials.Password == "" {
		http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

		return
	}

	jwt, err := authenticateService.Authenticate(credentials)
	if err != nil {
		if strings.Contains(err.Error(), u.ErrNotFound.Error()) {
			http.Error(w, u.ErrNotFound.Error(), http.StatusNotFound)

			return
		}
		if strings.Contains(err.Error(), u.ErrInvalidPassword.Error()) {
			http.Error(w, u.ErrInvalidPassword.Error(), http.StatusBadRequest)

			return
		}

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	io.WriteString(w, jwt)
}

func authorizeUser(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusBadRequest)

		return

	}

	jwt := strings.Split(authHeader, " ")[1]
	userID := authorizeService.Authorize(jwt)
	if userID == "" {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	io.WriteString(w, userID)
}
