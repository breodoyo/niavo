package router

import (
	"github.com/breodoyo/niavo/backend/internal/middleware"
	"github.com/breodoyo/niavo/backend/internal/app"
	"github.com/go-chi/chi/v5"
)

// New creates and configures the application's router.
func New() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
//Public route	
	router.Get("/", app.HomeHandler)

//API routes	
	router.Route("/api", func(r chi.Router) {

		r.Route("/v1", func(r chi.Router) {

			r.Get("/", app.HomeHandler)

			r.Route("/users", func(r chi.Router) {

			})
			r.Route("/organizations", func(r chi.Router) {

			})
			r.Route("/workflows", func(r chi.Router) {

			})
			r.Route("/workitems", func(r chi.Router) {

			})
		})
	})

	return router
}