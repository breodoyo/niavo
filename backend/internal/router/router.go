package router

import (
	"github.com/breodoyo/niavo/backend/internal/app"
	"github.com/breodoyo/niavo/backend/internal/middleware"
	"github.com/breodoyo/niavo/backend/internal/organization"
	"github.com/breodoyo/niavo/backend/internal/user"
	"github.com/go-chi/chi/v5"
)

// New creates and configures the application's router.
func New(orgHandler *organization.Handler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	//Public route
	router.Get("/", app.HomeHandler)

	//API routes
	router.Route("/api", func(r chi.Router) {

		r.Route("/v1", func(r chi.Router) {

			r.Get("/", app.HomeHandler)
			r.Get("/users", user.ListUsers)

			r.Route("/organizations", func(r chi.Router) {
				r.Post("/", orgHandler.CreateOrganization)
				r.Get("/", orgHandler.ListOrganizations)
				r.Get("/{id}", orgHandler.GetOrganization)
			})
			r.Route("/workflows", func(r chi.Router) {

			})
			r.Route("/workitems", func(r chi.Router) {

			})
		})
	})

	return router
}
