package user

import (
	"encoding/json"
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

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.WriteError(w, http.StatusBadRequest, "invalid_request", "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		common.WriteError(w, http.StatusBadRequest, "validation_error", err.Error())
		return
	}

	createdUser, err := h.service.CreateUser(r.Context(), req)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, "validation_error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createdUser); err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", "failed to encode response")
	}
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers(r.Context())
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", "failed to encode response")
	}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	u, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		common.WriteError(w, http.StatusNotFound, "not_found", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(u); err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", "failed to encode response")
	}
}