package main

import (
	"log"
	"net/http"

	"github.com/breodoyo/niavo/backend/internal/auth"
	"github.com/breodoyo/niavo/backend/internal/config"
	"github.com/breodoyo/niavo/backend/internal/database"
	"github.com/breodoyo/niavo/backend/internal/organization"
	"github.com/breodoyo/niavo/backend/internal/router"
	"github.com/breodoyo/niavo/backend/internal/user"
	"github.com/breodoyo/niavo/backend/internal/workflow"
	"github.com/breodoyo/niavo/backend/internal/workitem"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPool(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Printf(
		"%s (%s) listening on %s",
		cfg.AppName,
		cfg.Env,
		cfg.Port,
	)
	orgRepo := organization.NewRepository(db)
	orgService := organization.NewService(orgRepo)
	orgHandler := organization.NewHandler(orgService)

	UserRepo := user.NewRepository(db)
	UserService := user.NewService(UserRepo)
	UserHandler := user.NewHandler(UserService)

	workflowRepo := workflow.NewRepository(db)
	workflowService := workflow.NewService(workflowRepo)
	workflowHandler := workflow.NewHandler(workflowService)

	workitemRepo := workitem.NewRepository(db)
	workitemService := workitem.NewService(workitemRepo)
	workitemHandler := workitem.NewHandler(workitemService)

	authService := auth.NewService(UserService, cfg.JWTSecret)
	authHandler := auth.NewHandler(authService)

	r := router.New(orgHandler, UserHandler, authHandler, cfg.JWTSecret, workflowHandler, workitemHandler)

	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
