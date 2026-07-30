package organization

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateOrganization_InvalidRequest(t *testing.T) {
	handler := &Handler{
		service: nil,
	}

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

	recorder := httptest.NewRecorder()

	handler.CreateOrganization(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			recorder.Code,
		)
	}
}