package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"lewislbr/plantdex/plants/plant"

	"github.com/go-chi/chi/v5"
)

type (
	plantService interface {
		Add(string, plant.Plant) error
		ListAll(string) ([]plant.Plant, error)
		ListOne(string, string) (plant.Plant, error)
		Edit(string, string, plant.Plant) error
		Delete(string, string) error
	}

	handler struct {
		plantSvc plantService
	}
)

func NewHandler(plantSvc plantService) *handler {
	return &handler{plantSvc}
}

func (h *handler) Add(w http.ResponseWriter, r *http.Request) {
	var new plant.Plant

	err := json.NewDecoder(r.Body).Decode(&new)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)

		log.Printf("Error decoding add request: %v\n", err)

		return
	}

	userID := r.Context().Value(contextId).(string)
	err = h.plantSvc.Add(userID, new)
	if err != nil {
		switch {
		case errors.Is(err, plant.ErrMissingData):
			http.Error(w, plant.ErrMissingData.Error(), http.StatusBadRequest)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error adding plant: %v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) ListAll(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(contextId).(string)
	result, err := h.plantSvc.ListAll(userID)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)

		log.Printf("Error listing plants: %v\n", err)

		return
	}

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)

		log.Printf("Error encoding plants response: %v\n", err)

		return
	}
}

func (h *handler) ListOne(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(contextId).(string)
	plantID := chi.URLParam(r, "id")
	result, err := h.plantSvc.ListOne(userID, plantID)
	if err != nil {
		switch {
		case errors.Is(err, plant.ErrMissingData):
			http.Error(w, plant.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, plant.ErrNotFound):
			http.Error(w, plant.ErrNotFound.Error(), http.StatusNotFound)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error listing plant: %v\n", err)

			return
		}
	}

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)

		log.Printf("Error encoding plant response: %v\n", err)

		return
	}
}

func (h *handler) Edit(w http.ResponseWriter, r *http.Request) {
	var update plant.Plant

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)

		log.Printf("Error decoding edit request: %v\n", err)

		return
	}

	userID := r.Context().Value(contextId).(string)
	plantID := chi.URLParam(r, "id")
	err = h.plantSvc.Edit(userID, plantID, update)
	if err != nil {
		switch {
		case errors.Is(err, plant.ErrMissingData):
			http.Error(w, plant.ErrMissingData.Error(), http.StatusBadRequest)

			return
		case errors.Is(err, plant.ErrNotFound):
			http.Error(w, plant.ErrNotFound.Error(), http.StatusNotFound)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error editing plant: %v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(contextId).(string)
	plantID := chi.URLParam(r, "id")
	err := h.plantSvc.Delete(userID, plantID)
	if err != nil {
		switch {
		case errors.Is(err, plant.ErrMissingData):
			http.Error(w, plant.ErrMissingData.Error(), http.StatusBadRequest)

			return
		default:
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

			log.Printf("Error deleting plant: %v\n", err)

			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
