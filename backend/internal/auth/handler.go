package auth

import (
	"encoding/json"
	"net/http"

	"github.com/breodoyo/niavo/backend/internal/common"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.WriteError(w, http.StatusBadRequest, "invalid_request", "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		common.WriteError(w, http.StatusBadRequest, "validation_error", err.Error())
		return
	}

	token, err := h.service.Login(r.Context(), req)
	if err != nil {
		common.WriteError(w, http.StatusUnauthorized, "invalid_credentials", err.Error())
		return
	}

	resp := LoginResponse{Token: token}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		common.WriteError(w, http.StatusInternalServerError, "internal_error", "failed to encode response")
	}
}