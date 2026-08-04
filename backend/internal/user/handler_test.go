package user

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestCreateUser_InvalidRequest(t *testing.T) {
	handler := NewHandler(&MockUserService{})

	body := []byte(`{
		"email": "lben@gmailcom",
		"first_name": "Ben",
		"last_name": "Luis"
		"password": "Hello123!"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/users",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.CreateUser(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}
func TestCreateUser_Success(t *testing.T) {
	handler := NewHandler(&MockUserService{})

	body := []byte(`{
		"email": "lben@gmail.com",
		"first_name": "Ben",
		"last_name": "Luis",
		"password": "Hello123!"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/users",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.CreateUser(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusCreated,
			rec.Code,
		)
	}
}
func TestUpdateUser_Success(t *testing.T) {
	handler := NewHandler(&MockUserService{})

	body := []byte(`{
		"first_name": "Bree",
		"last_name": "Odoyo"
	}`)

	req := httptest.NewRequest(
		http.MethodPatch,
		"/users/123",
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "123")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rec := httptest.NewRecorder()

	handler.UpdateUser(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestDeleteUser_Success(t *testing.T) {
	handler := NewHandler(&MockUserService{})

	req := httptest.NewRequest(http.MethodDelete, "/users/123", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "123")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rec := httptest.NewRecorder()

	handler.DeleteUser(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}
}
