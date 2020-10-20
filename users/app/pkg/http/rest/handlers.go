package rest

import (
	"users/pkg/add"
	"users/pkg/delete"
	"users/pkg/edit"
	"users/pkg/entity"
	"users/pkg/list"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getUser(s list.Service) func(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")

		user, err := s.GetUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("content-type", "application/json")

		json.NewEncoder(w).Encode(user)
	}
}

func addUser(s add.Service) func(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var user entity.User

		json.NewDecoder(r.Body).Decode(&user)

		s.AddUser(user)

		w.Write([]byte(user.Name + " added"))
	}
}

func editUser(s edit.Service) func(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")

		var user entity.User

		json.NewDecoder(r.Body).Decode(&user)

		err := s.EditUser(id, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Write([]byte(user.Name + " updated"))
	}
}

func deleteUser(s delete.Service) func(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")

		err := s.DeleteUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Write([]byte(id + " deleted"))
	}
}
