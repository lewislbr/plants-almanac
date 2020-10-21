package rest

import (
	"io"
	"users/pkg/edit"
	u "users/pkg/user"

	"encoding/json"
	"net/http"
	"users/pkg/add"
	"users/pkg/delete"
	"users/pkg/list"

	"github.com/julienschmidt/httprouter"
)

func getUser(ls list.Service) func(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		user, err := ls.ListUser(u.ID(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("content-type", "application/json")

		json.NewEncoder(w).Encode(user)
	}
}

func addUser(as add.Service) func(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var user u.User

		json.NewDecoder(r.Body).Decode(&user)

		as.AddUser(user)

		io.WriteString(w, user.Name+" added")
	}
}

func editUser(es edit.Service) func(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")

		var user u.User

		json.NewDecoder(r.Body).Decode(&user)

		err := es.EditUser(u.ID(id), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		io.WriteString(w, user.Name+" updated")
	}
}

func deleteUser(ds delete.Service) func(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		err := ds.DeleteUser(u.ID(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		io.WriteString(w, id+" deleted")
	}
}
