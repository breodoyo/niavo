package organization

import (
	"encoding/json"
	"net/http"
	"errors"

	"github.com/breodoyo/niavo/backend/internal/common"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service OrganizationService
}

func NewHandler(service OrganizationService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	var req CreateOrganizationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.WriteError(
			w,
			http.StatusBadRequest,
			"invalid_request",
			"invalid request body",
		)
		return
	}

	if err := req.Validate(); err != nil {
		common.WriteError(
			w,
			http.StatusBadRequest,
			"validation_error",
			err.Error(),
		)
		return
	}

	org := Organization{
		Name: req.Name,
		Slug: req.Slug,
	}

	createOrg, err := h.service.CreateOrganization(r.Context(), org)
	if err != nil {
		common.WriteError(
			w,
			http.StatusBadRequest,
			"validation_error",
			err.Error(),
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createOrg); err != nil {
		common.WriteError(
			w,
			http.StatusInternalServerError,
			"internal_error",
			"failed to encode response",
		)
	}
}

func (h *Handler) ListOrganizations(w http.ResponseWriter, r *http.Request) {
	organizations, err := h.service.ListOrganizations(r.Context())
	if err != nil {
		common.WriteError(
			w,
			http.StatusInternalServerError,
			"internal_error",
			err.Error(),
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(organizations); err != nil {
		common.WriteError(
			w,
			http.StatusInternalServerError,
			"internal_error",
			"failed to encode response",
		)
	}
}

func (h *Handler) GetOrganization(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	org, err := h.service.GetOrganization(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrOrganizationNotFound) {
			common.WriteError(w, http.StatusNotFound, "not_found", "organization not found")
			return
		}
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(org); err != nil {
		common.WriteError(
			w,
			http.StatusInternalServerError,
			"internal_error",
			"failed to encode response",
		)
	}
}

func (h *Handler) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req UpdateOrganizationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.WriteError(w, http.StatusBadRequest, "invalid_request", "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		common.WriteError(w, http.StatusBadRequest, "validation_error", err.Error())
		return
	}

	updateOrg, err := h.service.UpdateOrganization(r.Context(), id, req.Name)
	if err != nil {
		if errors.Is(err, ErrOrganizationNotFound) {
			common.WriteError(w, http.StatusNotFound, "not_found", "organization not found")
			return
		}
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(updateOrg); err != nil {
		common.WriteError(
			w,
			http.StatusInternalServerError,
			"internal_error",
			"failed to encode response",
		)
	}
}

func (h *Handler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.service.DeleteOrganization(r.Context(), id); err != nil {
		if errors.Is(err, ErrOrganizationNotFound) {
			common.WriteError(w, http.StatusNotFound, "not_found", "organization not found")
			return
		}
		common.WriteError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}