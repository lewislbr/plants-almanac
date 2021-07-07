package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"plants/plant"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	t.Run("should return 201 after successful request", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("Add", mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(nil)

		plant := &plant.Plant{Name: "test"}
		payload, err := json.Marshal(plant)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.Add(w, r.WithContext(ctx))

		require.Equal(t, http.StatusCreated, w.Result().StatusCode)
	})

	t.Run("should return 400 if required data is missing", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("Add", mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(plant.ErrMissingData)

		var plant plant.Plant

		payload, err := json.Marshal(plant)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.Add(w, r.WithContext(ctx))

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("Add", mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(errors.New("error"))

		plant := &plant.Plant{Name: "test"}
		payload, err := json.Marshal(plant)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.Add(w, r.WithContext(ctx))

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestListAll(t *testing.T) {
	t.Run("should return 200 if the request is successful", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("ListAll", mock.AnythingOfType("string")).Return([]plant.Plant{}, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.ListAll(w, r.WithContext(ctx))

		require.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("should return 500 if if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("ListAll", mock.AnythingOfType("string")).Return(nil, errors.New("error"))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.ListAll(w, r.WithContext(ctx))

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestListOne(t *testing.T) {
	t.Run("should return 200 if the request is successful", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("ListOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{Name: "test"}, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/123", nil)
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.ListOne(w, r.WithContext(ctx))

		require.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("should return 400 if required data is missing", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("ListOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, plant.ErrMissingData)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.ListOne(w, r.WithContext(ctx))

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("should return 404 if the plant is not found", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("ListOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, plant.ErrNotFound)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/123", nil)
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.ListOne(w, r.WithContext(ctx))

		require.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("ListOne", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.Plant{}, errors.New("error"))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/123", nil)
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.ListOne(w, r.WithContext(ctx))

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestEdit(t *testing.T) {
	t.Run("should return 204 if the request is successful", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("Edit", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(nil)

		var plant plant.Plant

		payload, err := json.Marshal(plant)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/123", bytes.NewBuffer(payload))
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.Edit(w, r.WithContext(ctx))

		require.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("should return 400 if required data is missing", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("Edit", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(plant.ErrMissingData)

		var plant plant.Plant

		payload, err := json.Marshal(plant)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(payload))
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.Edit(w, r.WithContext(ctx))

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("should return 404 if the plant is not found", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("Edit", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(plant.ErrNotFound)

		var plant plant.Plant

		payload, err := json.Marshal(plant)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/123", bytes.NewBuffer(payload))
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.Edit(w, r.WithContext(ctx))

		require.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("Edit", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("plant.Plant")).Return(errors.New("error"))

		var plant plant.Plant

		payload, err := json.Marshal(plant)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/123", bytes.NewBuffer(payload))
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.Edit(w, r.WithContext(ctx))

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should return 204 if the request is successful", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("Delete", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/123", nil)
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.Delete(w, r.WithContext(ctx))

		require.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("should return 400 if required data is missing", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("Delete", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(plant.ErrMissingData)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/", nil)
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.Delete(w, r.WithContext(ctx))

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		plantSvc := &mockPlantService{}
		handler := NewHandler(plantSvc)

		plantSvc.On("Delete", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(errors.New("error"))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/123", nil)
		ctx := context.WithValue(r.Context(), contextId, "abc")

		handler.Delete(w, r.WithContext(ctx))

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
