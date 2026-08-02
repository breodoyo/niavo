package user

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
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
