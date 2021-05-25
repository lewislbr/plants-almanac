package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"users/token"
	"users/user"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Run("should return 201 after successful request", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		userSvc.On("Create", mock.AnythingOfType("user.User")).Return(nil)

		user := &user.User{Name: "test", Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))

		handler.Create(w, r)

		require.Equal(t, http.StatusCreated, w.Result().StatusCode)
	})

	t.Run("should return 400 if required data is missing", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		userSvc.On("Create", mock.AnythingOfType("user.User")).Return(user.ErrMissingData)

		user := &user.User{Name: "test", Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))

		handler.Create(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("should return 409 if a user already exists", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		userSvc.On("Create", mock.AnythingOfType("user.User")).Return(user.ErrUserExists)

		user := &user.User{Name: "test", Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))

		handler.Create(w, r)

		require.Equal(t, http.StatusConflict, w.Result().StatusCode)
	})

	t.Run("should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		userSvc.On("Create", mock.AnythingOfType("user.User")).Return(errors.New("error"))

		user := &user.User{Name: "test", Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))

		handler.Create(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
func TestLogIn(t *testing.T) {
	t.Run("should return 204 if the request is successful", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		userSvc.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", nil)
		tokenSvc.On("Generate", mock.AnythingOfType("string")).Return("", nil)

		user := &user.Credentials{Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))

		handler.LogIn(w, r)

		require.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("should return 400 if required data is missing", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		userSvc.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", user.ErrMissingData)
		tokenSvc.On("Generate", mock.AnythingOfType("string")).Return("", user.ErrMissingData)

		user := &user.Credentials{Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))

		handler.LogIn(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("should return 404 if the user is not found", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		userSvc.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", user.ErrNotFound)
		tokenSvc.On("Generate", mock.AnythingOfType("string")).Return("", nil)

		user := &user.Credentials{Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))

		handler.LogIn(w, r)

		require.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("should return 400 if the password is invalid", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		userSvc.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", user.ErrInvalidPassword)
		tokenSvc.On("Generate", mock.AnythingOfType("string")).Return("", nil)

		user := &user.Credentials{Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))

		handler.LogIn(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		userSvc.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", errors.New("error"))
		tokenSvc.On("Generate", mock.AnythingOfType("string")).Return("", errors.New("error"))

		user := &user.Credentials{Email: "test", Password: "test"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))

		handler.LogIn(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
func TestAuthorize(t *testing.T) {
	t.Run("should return 200 if the request is successful", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", nil)

		handler.Authorize(w, r)

		require.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("should return 400 if the token is empty", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", token.ErrMissingData)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", nil)

		handler.Authorize(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("should return 401 if the token is invalid", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", token.ErrInvalidToken)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", nil)

		handler.Authorize(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", errors.New("error"))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", nil)

		r.Header.Add("Authorization", "Bearer test")

		handler.Authorize(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestRefresh(t *testing.T) {
	t.Run("should return 204 if the request is successful", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", nil)
		tokenSvc.On("Revoke", mock.AnythingOfType("string")).Return(nil)
		tokenSvc.On("Generate", mock.AnythingOfType("string")).Return("", nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Refresh(w, r)

		require.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("should return 400 if the token is missing", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", token.ErrMissingData)
		tokenSvc.On("Revoke", mock.AnythingOfType("string")).Return(token.ErrMissingData)
		tokenSvc.On("Generate", mock.AnythingOfType("string")).Return("", token.ErrMissingData)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Refresh(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("should return 401 if the token is invalid", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", token.ErrInvalidToken)
		tokenSvc.On("Revoke", mock.AnythingOfType("string")).Return(token.ErrInvalidToken)
		tokenSvc.On("Generate", mock.AnythingOfType("string")).Return("", nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Refresh(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", errors.New("error"))
		tokenSvc.On("Revoke", mock.AnythingOfType("string")).Return(errors.New("error"))
		tokenSvc.On("Generate", mock.AnythingOfType("string")).Return("", errors.New("error"))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Refresh(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestLogOut(t *testing.T) {
	t.Run("should return 204 if the request is successful", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Revoke", mock.AnythingOfType("string")).Return(nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.LogOut(w, r)

		require.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("should return 400 if the token is missing", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Revoke", mock.AnythingOfType("string")).Return(token.ErrMissingData)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.LogOut(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("should return 401 if the token is invalid", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Revoke", mock.AnythingOfType("string")).Return(token.ErrInvalidToken)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.LogOut(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Revoke", mock.AnythingOfType("string")).Return(errors.New("error"))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.LogOut(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestInfo(t *testing.T) {
	t.Run("should return 200 if the request is successful", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", nil)
		userSvc.On("Info", mock.AnythingOfType("string")).Return(user.Info{}, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Info(w, r)

		require.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("should return 400 if the user ID is missing", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", token.ErrMissingData)
		userSvc.On("Info", mock.AnythingOfType("string")).Return(user.Info{}, user.ErrMissingData)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Info(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("should return 401 if the token is invalid", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", token.ErrInvalidToken)
		userSvc.On("Info", mock.AnythingOfType("string")).Return(user.Info{}, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Info(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("should return 404 if the user is not found", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", nil)
		userSvc.On("Info", mock.AnythingOfType("string")).Return(user.Info{}, user.ErrNotFound)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Info(w, r)

		require.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		userSvc := &mockUserService{}
		tokenSvc := &mockTokenService{}
		handler := NewHandler(userSvc, tokenSvc, "")

		tokenSvc.On("Validate", mock.AnythingOfType("string")).Return("", errors.New("error"))
		userSvc.On("Info", mock.AnythingOfType("string")).Return(user.Info{}, errors.New("error"))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Info(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
