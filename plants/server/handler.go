package server

import (
	"encoding/json"
	"log"
	"net/http"

	"plants/plant"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
)

type (
	adder interface {
		Add(string, plant.Plant) error
	}
	lister interface {
		ListAll(string) ([]plant.Plant, error)
		ListOne(string, string) (plant.Plant, error)
	}
	editer interface {
		Edit(string, string, plant.Plant) error
	}
	deleter interface {
		Delete(string, string) error
	}

	handler struct {
		addSvc    adder
		listSvc   lister
		editSvc   editer
		deleteSvc deleter
	}
)

func NewHandler(addSvc adder, listSvc lister, editSvc editer, deleteSvc deleter) *handler {
	return &handler{addSvc, listSvc, editSvc, deleteSvc}
}

func (h *handler) Add(w http.ResponseWriter, r *http.Request) {
	new := plant.Plant{}
	err := json.NewDecoder(r.Body).Decode(&new)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	userID := r.Context().Value(contextId).(string)
	err = h.addSvc.Add(userID, new)
	if err != nil {
		switch {
		case errors.Is(err, plant.ErrMissingData):
			http.Error(w, plant.ErrMissingData.Error(), http.StatusBadRequest)

			return
		default:
			http.Error(w, "something went wrong", http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) ListAll(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(contextId).(string)
	result, err := h.listSvc.ListAll(userID)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}
}

func (h *handler) ListOne(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(contextId).(string)
	plantID := chi.URLParam(r, "id")
	result, err := h.listSvc.ListOne(userID, plantID)
	if err != nil {
		switch {
		case errors.Is(err, plant.ErrMissingData):
			http.Error(w, plant.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, plant.ErrNotFound):
			http.Error(w, plant.ErrNotFound.Error(), http.StatusNotFound)

			return
		default:
			http.Error(w, "something went wrong", http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}
}

func (h *handler) Edit(w http.ResponseWriter, r *http.Request) {
	update := plant.Plant{}
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)

		log.Printf("%+v\n", err)

		return
	}

	userID := r.Context().Value(contextId).(string)
	plantID := chi.URLParam(r, "id")
	err = h.editSvc.Edit(userID, plantID, update)
	if err != nil {
		switch {
		case errors.Is(err, plant.ErrMissingData):
			http.Error(w, plant.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, plant.ErrNotFound):
			http.Error(w, plant.ErrNotFound.Error(), http.StatusNotFound)

			return
		default:
			http.Error(w, "something went wrong", http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(contextId).(string)
	plantID := chi.URLParam(r, "id")
	err := h.deleteSvc.Delete(userID, plantID)
	if err != nil {
		switch {
		case errors.Is(err, plant.ErrMissingData):
			http.Error(w, plant.ErrMissingData.Error(), http.StatusBadRequest)

			return
		default:
			http.Error(w, "something went wrong", http.StatusInternalServerError)

			log.Printf("%+v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
