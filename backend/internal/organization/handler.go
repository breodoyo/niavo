package organization

import (
	"encoding/json"

	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}
type UpdateOrganizationRequest struct {
	Name string `json:"name"`
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}
func (h *Handler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	var org Organization

	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	createOrg, err := h.service.CreateOrganization(r.Context(), org)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createOrg); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
func (h *Handler) ListOrganizations(w http.ResponseWriter, r *http.Request) {

	organizations, err := h.service.ListOrganizations(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(organizations); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
func (h *Handler) GetOrganization(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	org, err := h.service.GetOrganization(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(org); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
func (h *Handler) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req UpdateOrganizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	updateOrg, err := h.service.UpdateOrganization(r.Context(), id, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(updateOrg); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
func (h *Handler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.service.DeleteOrganization(r.Context(), id); err != nil {
		http.Error(w, "Organization not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
