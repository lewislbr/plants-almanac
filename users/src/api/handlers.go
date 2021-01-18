package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"users/src/authenticate"
	"users/src/authorize"
	"users/src/create"
	"users/src/storage"
	u "users/src/user"
)

var (
	createService       = create.NewService(&storage.MongoDB{})
	authenticateService = authenticate.NewService(&storage.MongoDB{})
	authorizeService    = authorize.NewService()
)

func createUser(w http.ResponseWriter, r *http.Request) {
	var new u.User

	json.NewDecoder(r.Body).Decode(&new)

	if new.Name == "" || new.Email == "" || new.Password == "" {
		http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

		return
	}

	err := createService.Create(new)
	if err != nil {
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

func logInUser(w http.ResponseWriter, r *http.Request) {
	var cred u.Credentials

	json.NewDecoder(r.Body).Decode(&cred)

	if cred.Email == "" || cred.Password == "" {
		http.Error(w, u.ErrMissingData.Error(), http.StatusBadRequest)

		return
	}

	jwt, err := authenticateService.Authenticate(cred)
	if err != nil {
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

func authorizeUser(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	jwt := strings.Split(authHeader, " ")[1]
	userID, err := authorizeService.Authorize(jwt)
	if err != nil {
		if err == u.ErrInvalidToken {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}

		w.WriteHeader(http.StatusInternalServerError)

		log.Printf("%+v\n", err)
	}

	io.WriteString(w, userID)
}
