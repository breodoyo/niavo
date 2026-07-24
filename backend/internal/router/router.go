package router

import (
	"github.com/breodoyo/niavo/backend/internal/app"
	"github.com/go-chi/chi/v5"
)

// New creates and configures the application's router.
func New() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", app.HomeHandler)

	return router
}