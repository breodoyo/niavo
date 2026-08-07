package workflow

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestCreateWorkflow_InvalidRequest(t *testing.T) {
	handler := NewHandler(&MockWorkflowService{})

	body := []byte(`{"name": "", "description": "test"}`)
	req := httptest.NewRequest(http.MethodPost, "/workflows", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler.CreateWorkflow(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestCreateWorkflow_Success(t *testing.T) {
	handler := NewHandler(&MockWorkflowService{})

	body := []byte(`{"name": "Support Tickets", "description": "Track requests"}`)
	req := httptest.NewRequest(http.MethodPost, "/workflows", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler.CreateWorkflow(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}
}

func TestUpdateWorkflow_Success(t *testing.T) {
	handler := NewHandler(&MockWorkflowService{})

	body := []byte(`{"name": "Updated Name", "description": "Updated desc"}`)
	req := httptest.NewRequest(http.MethodPatch, "/workflows/123", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "123")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rec := httptest.NewRecorder()
	handler.UpdateWorkflow(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestDeleteWorkflow_Success(t *testing.T) {
	handler := NewHandler(&MockWorkflowService{})

	req := httptest.NewRequest(http.MethodDelete, "/workflows/123", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "123")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rec := httptest.NewRecorder()
	handler.DeleteWorkflow(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}
}