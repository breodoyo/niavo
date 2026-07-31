package organization

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateOrganization_InvalidRequest(t *testing.T) {
	handler := NewHandler(&MockOrganizationService{})

	body := []byte(`{
		"name": "",
		"slug": "niavo"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/organizations",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.CreateOrganization(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}
func TestCreateOrganization_Success(t *testing.T) {
	handler := NewHandler(&MockOrganizationService{})

	body := []byte(`{
		"name": "Niavo",
		"slug": "niavo"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/organizations",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.CreateOrganization(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusCreated,
			rec.Code,
		)
	}
}
