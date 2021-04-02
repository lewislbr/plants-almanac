package server

import (
	"encoding/json"
	"log"
	"net/http"

	"plants/add"
	"plants/delete"
	"plants/edit"
	"plants/list"
	"plants/plant"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

type handler struct {
	as add.AddService
	ls list.ListService
	es edit.EditService
	ds delete.DeleteService
}

func NewHandler(as add.AddService, ls list.ListService, es edit.EditService, ds delete.DeleteService) *handler {
	return &handler{as, ls, es, ds}
}

func (h *handler) Add(w http.ResponseWriter, r *http.Request) {
	new := plant.Plant{}
	err := json.NewDecoder(r.Body).Decode(&new)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	uid := r.Context().Value(contextId).(string)
	err = h.as.Add(uid, new)
	if err != nil {
		switch {
		case errors.Is(err, plant.ErrMissingData):
			http.Error(w, plant.ErrMissingData.Error(), http.StatusBadRequest)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) ListAll(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(contextId).(string)
	result, err := h.ls.ListAll(uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}
}

func (h *handler) ListOne(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(contextId).(string)
	id := chi.URLParam(r, "id")
	result, err := h.ls.ListOne(uid, id)
	if err != nil {
		switch {
		case errors.Is(err, plant.ErrMissingData):
			http.Error(w, plant.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, plant.ErrNotFound):
			http.Error(w, plant.ErrNotFound.Error(), http.StatusNotFound)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}
}

func (h *handler) Edit(w http.ResponseWriter, r *http.Request) {
	update := plant.Plant{}
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	uid := r.Context().Value(contextId).(string)
	id := chi.URLParam(r, "id")
	err = h.es.Edit(uid, id, update)
	if err != nil {
		switch {
		case errors.Is(err, plant.ErrMissingData):
			http.Error(w, plant.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, plant.ErrNotFound):
			http.Error(w, plant.ErrNotFound.Error(), http.StatusNotFound)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(contextId).(string)
	id := chi.URLParam(r, "id")
	err := h.ds.Delete(uid, id)
	if err != nil {
		switch {
		case errors.Is(err, plant.ErrMissingData):
			http.Error(w, plant.ErrMissingData.Error(), http.StatusBadRequest)

			return
		default:
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
