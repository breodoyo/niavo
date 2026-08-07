package workflow

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/breodoyo/niavo/backend/internal/common"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateWorkflow(w http.ResponseWriter, r *http.Request) {
	var req CreateWorkflowRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.WriteError(w, http.StatusBadRequest, "invalid_request", "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		common.WriteError(w, http.StatusBadRequest, "validation_error", err.Error())
		return
	}

	wf, err := h.service.CreateWorkflow(r.Context(), req)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(wf); err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", "failed to encode response")
	}
}

func (h *Handler) ListWorkflows(w http.ResponseWriter, r *http.Request) {
	workflows, err := h.service.ListWorkflows(r.Context())
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(workflows); err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", "failed to encode response")
	}
}

func (h *Handler) GetWorkflow(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	wf, err := h.service.GetWorkflow(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrWorkflowNotFound) {
			common.WriteError(w, http.StatusNotFound, "not_found", "workflow not found")
			return
		}
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(wf); err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", "failed to encode response")
	}
}

func (h *Handler) UpdateWorkflow(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req UpdateWorkflowRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.WriteError(w, http.StatusBadRequest, "invalid_request", "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		common.WriteError(w, http.StatusBadRequest, "validation_error", err.Error())
		return
	}

	wf, err := h.service.UpdateWorkflow(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, ErrWorkflowNotFound) {
			common.WriteError(w, http.StatusNotFound, "not_found", "workflow not found")
			return
		}
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(wf); err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", "failed to encode response")
	}
}

func (h *Handler) DeleteWorkflow(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.service.DeleteWorkflow(r.Context(), id); err != nil {
		if errors.Is(err, ErrWorkflowNotFound) {
			common.WriteError(w, http.StatusNotFound, "not_found", "workflow not found")
			return
		}
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}