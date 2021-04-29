package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"users/user"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	t.Run("Create should return 201 after successful request", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		creater.On("Create", mock.AnythingOfType("user.User")).Return(nil)

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

	t.Run("Create should return 400 if required data is missing", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		creater.On("Create", mock.AnythingOfType("user.User")).Return(user.ErrMissingData)

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

	t.Run("Create should return 409 if a user already exists", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		creater.On("Create", mock.AnythingOfType("user.User")).Return(user.ErrUserExists)

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

	t.Run("Create should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		creater.On("Create", mock.AnythingOfType("user.User")).Return(errors.New("error"))

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

	t.Run("LogIn should return 204 if the request is successful", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authenticater.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", nil)
		generater.On("GenerateToken", mock.AnythingOfType("string")).Return("", nil)

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

	t.Run("LogIn should return 400 if required data is missing", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authenticater.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", user.ErrMissingData)
		generater.On("GenerateToken", mock.AnythingOfType("string")).Return("", user.ErrMissingData)

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

	t.Run("LogIn should return 404 if the user is not found", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authenticater.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", user.ErrNotFound)
		generater.On("GenerateToken", mock.AnythingOfType("string")).Return("", nil)

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

	t.Run("LogIn should return 400 if the password is invalid", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authenticater.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", user.ErrInvalidPassword)
		generater.On("GenerateToken", mock.AnythingOfType("string")).Return("", nil)

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

	t.Run("LogIn should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authenticater.On("Authenticate", mock.AnythingOfType("user.Credentials")).Return("", errors.New("error"))
		generater.On("GenerateToken", mock.AnythingOfType("string")).Return("", errors.New("error"))

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

	t.Run("Authorize should return 200 if the request is successful", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", nil)

		handler.Authorize(w, r)

		require.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("Authorize should return 400 if the token is empty", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", user.ErrMissingData)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", nil)

		handler.Authorize(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("Authorize should return 401 if the token is invalid", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", user.ErrInvalidToken)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", nil)

		handler.Authorize(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("Authorize should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", errors.New("error"))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", nil)

		r.Header.Add("Authorization", "Bearer test")

		handler.Authorize(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Refresh should return 204 if the request is successful", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", nil)
		revoker.On("RevokeToken", mock.AnythingOfType("string")).Return(nil)
		generater.On("GenerateToken", mock.AnythingOfType("string")).Return("", nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Refresh(w, r)

		require.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("Refresh should return 400 if the token is missing", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", user.ErrMissingData)
		revoker.On("RevokeToken", mock.AnythingOfType("string")).Return(user.ErrMissingData)
		generater.On("GenerateToken", mock.AnythingOfType("string")).Return("", user.ErrMissingData)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Refresh(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("Refresh should return 401 if the token is invalid", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", user.ErrInvalidToken)
		revoker.On("RevokeToken", mock.AnythingOfType("string")).Return(user.ErrInvalidToken)
		generater.On("GenerateToken", mock.AnythingOfType("string")).Return("", nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Refresh(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("Refresh should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", errors.New("error"))
		revoker.On("RevokeToken", mock.AnythingOfType("string")).Return(errors.New("error"))
		generater.On("GenerateToken", mock.AnythingOfType("string")).Return("", errors.New("error"))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Refresh(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("LogOut should return 204 if the request is successful", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		revoker.On("RevokeToken", mock.AnythingOfType("string")).Return(nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.LogOut(w, r)

		require.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("LogOut should return 400 if the token is missing", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		revoker.On("RevokeToken", mock.AnythingOfType("string")).Return(user.ErrMissingData)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.LogOut(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("LogOut should return 401 if the token is invalid", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		revoker.On("RevokeToken", mock.AnythingOfType("string")).Return(user.ErrInvalidToken)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.LogOut(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("LogOut should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		revoker.On("RevokeToken", mock.AnythingOfType("string")).Return(errors.New("error"))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.LogOut(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Info should return 200 if the request is successful", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", nil)
		infoer.On("UserInfo", mock.AnythingOfType("string")).Return(user.Info{}, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Info(w, r)

		require.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("Info should return 400 if the user ID is missing", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", user.ErrMissingData)
		infoer.On("UserInfo", mock.AnythingOfType("string")).Return(user.Info{}, user.ErrMissingData)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Info(w, r)

		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("Info should return 401 if the token is invalid", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", user.ErrInvalidToken)
		infoer.On("UserInfo", mock.AnythingOfType("string")).Return(user.Info{}, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Info(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("Info should return 404 if the user is not found", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", nil)
		infoer.On("UserInfo", mock.AnythingOfType("string")).Return(user.Info{}, user.ErrNotFound)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Info(w, r)

		require.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("Info should return 500 if an unexpected error happens", func(t *testing.T) {
		t.Parallel()

		creater := &mockCreater{}
		authenticater := &mockAuthenticater{}
		authorizer := &mockAuthorizer{}
		generater := &mockGenerater{}
		revoker := &mockRevoker{}
		infoer := &mockInfoer{}
		handler := NewHandler(creater, authenticater, authorizer, generater, revoker, infoer, "")

		authorizer.On("Authorize", mock.AnythingOfType("string")).Return("", errors.New("error"))
		infoer.On("UserInfo", mock.AnythingOfType("string")).Return(user.Info{}, errors.New("error"))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.Info(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
