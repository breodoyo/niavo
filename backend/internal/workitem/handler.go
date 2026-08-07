package workitem

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
	return &Handler{service: service}
}

func (h *Handler) CreateWorkItem(w http.ResponseWriter, r *http.Request) {
	var req CreateWorkItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.WriteError(w, http.StatusBadRequest, "invalid_request", "invalid request body")
		return
	}
	if err := req.Validate(); err != nil {
		common.WriteError(w, http.StatusBadRequest, "validation_error", err.Error())
		return
	}
	item, err := h.service.CreateWorkItem(r.Context(), req)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) ListWorkItems(w http.ResponseWriter, r *http.Request) {
	workflowID := r.URL.Query().Get("workflow_id")
	items, err := h.service.ListWorkItems(r.Context(), workflowID)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) GetWorkItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	item, err := h.service.GetWorkItem(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrWorkItemNotFound) {
			common.WriteError(w, http.StatusNotFound, "not_found", "work item not found")
			return
		}
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) UpdateWorkItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req UpdateWorkItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.WriteError(w, http.StatusBadRequest, "invalid_request", "invalid request body")
		return
	}
	if err := req.Validate(); err != nil {
		common.WriteError(w, http.StatusBadRequest, "validation_error", err.Error())
		return
	}
	item, err := h.service.UpdateWorkItem(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, ErrWorkItemNotFound) {
			common.WriteError(w, http.StatusNotFound, "not_found", "work item not found")
			return
		}
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) DeleteWorkItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.service.DeleteWorkItem(r.Context(), id); err != nil {
		if errors.Is(err, ErrWorkItemNotFound) {
			common.WriteError(w, http.StatusNotFound, "not_found", "work item not found")
			return
		}
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}