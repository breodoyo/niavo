package router

import (
	"github.com/breodoyo/niavo/backend/internal/app"
	"github.com/breodoyo/niavo/backend/internal/middleware"
	"github.com/breodoyo/niavo/backend/internal/organization"
	"github.com/breodoyo/niavo/backend/internal/user"
	"github.com/go-chi/chi/v5"
)

// New creates and configures the application's router.
func New(orgHandler *organization.Handler, userHandler *user.Handler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	//Public route
	router.Get("/", app.HomeHandler)

	//API routes
	router.Route("/api", func(r chi.Router) {

		r.Route("/v1", func(r chi.Router) {

			r.Get("/", app.HomeHandler)

			r.Route("/users", func(r chi.Router) {
				r.Post("/", userHandler.CreateUser)
				r.Get("/", userHandler.ListUsers)
				r.Get("/{id}", userHandler.GetUser)
			})

			r.Route("/organizations", func(r chi.Router) {
				r.Post("/", orgHandler.CreateOrganization)
				r.Get("/", orgHandler.ListOrganizations)
				r.Get("/{id}", orgHandler.GetOrganization)
				r.Patch("/{id}", orgHandler.UpdateOrganization)
				r.Delete("/{id}", orgHandler.DeleteOrganization)
			})
			r.Route("/workflows", func(r chi.Router) {

			})
			r.Route("/workitems", func(r chi.Router) {

			})
		})
	})

	return router
}