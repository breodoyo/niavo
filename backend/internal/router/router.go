package router

import (
	"github.com/breodoyo/niavo/backend/internal/app"
	"github.com/breodoyo/niavo/backend/internal/auth"
	"github.com/breodoyo/niavo/backend/internal/middleware"
	"github.com/breodoyo/niavo/backend/internal/organization"
	"github.com/breodoyo/niavo/backend/internal/user"
	"github.com/breodoyo/niavo/backend/internal/workflow"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func New(orgHandler *organization.Handler, userHandler *user.Handler, authHandler *auth.Handler, jwtSecret string, workflowHandler *workflow.Handler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/", app.HomeHandler)

	router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {

			r.Get("/", app.HomeHandler)

			r.Route("/auth", func(r chi.Router) {
				r.Post("/login", authHandler.Login)
			})

			r.Route("/users", func(r chi.Router) {
				r.Post("/", userHandler.CreateUser)

				r.Group(func(r chi.Router) {
					r.Use(middleware.RequireAuth(jwtSecret))
					r.Get("/", userHandler.ListUsers)
					r.Get("/{id}", userHandler.GetUser)
					r.Patch("/{id}", userHandler.UpdateUser)
					r.Delete("/{id}", userHandler.DeleteUser)
				})
			})

			r.Route("/organizations", func(r chi.Router) {
				r.Use(middleware.RequireAuth(jwtSecret))
				r.Post("/", orgHandler.CreateOrganization)
				r.Get("/", orgHandler.ListOrganizations)
				r.Get("/{id}", orgHandler.GetOrganization)
				r.Patch("/{id}", orgHandler.UpdateOrganization)
				r.Delete("/{id}", orgHandler.DeleteOrganization)
			})

			r.Route("/workflows", func(r chi.Router) {
				r.Use(middleware.RequireAuth(jwtSecret))
				r.Post("/", workflowHandler.CreateWorkflow)
				r.Get("/", workflowHandler.ListWorkflows)
				r.Get("/{id}", workflowHandler.GetWorkflow)
				r.Patch("/{id}", workflowHandler.UpdateWorkflow)
				r.Delete("/{id}", workflowHandler.DeleteWorkflow)
			})
			r.Route("/workitems", func(r chi.Router) {})
		})
	})

	return router
}
