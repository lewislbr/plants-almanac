package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"users/pkg/add"
	"users/pkg/delete"
	"users/pkg/edit"
	"users/pkg/list"
	u "users/pkg/user"

	"github.com/julienschmidt/httprouter"
)

func addUser(as add.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user u.User

		json.NewDecoder(r.Body).Decode(&user)

		as.AddUser(user)

		io.WriteString(w, user.Name+" added")
	}
}

func listUser(ls list.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		id := params.ByName("id")
		user, err := ls.ListUser(u.ID(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("content-type", "application/json")

		json.NewEncoder(w).Encode(user)
	}
}

func editUser(es edit.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user u.User

		json.NewDecoder(r.Body).Decode(&user)

		params := httprouter.ParamsFromContext(r.Context())
		id := params.ByName("id")
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
) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		id := params.ByName("id")
		err := ds.DeleteUser(u.ID(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		io.WriteString(w, id+" deleted")
	}
}
