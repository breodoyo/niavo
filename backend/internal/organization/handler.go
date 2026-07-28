package organization

import (
	"encoding/json"

	"net/http"
)

type Handler struct {
	service *Service
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
