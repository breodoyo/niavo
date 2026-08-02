package main

import (
	"log"
	"net/http"

	"github.com/breodoyo/niavo/backend/internal/config"
	"github.com/breodoyo/niavo/backend/internal/database"
	"github.com/breodoyo/niavo/backend/internal/organization"
	"github.com/breodoyo/niavo/backend/internal/router"
	"github.com/breodoyo/niavo/backend/internal/user"
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

	userRepo := user.NewRepository(db)
	UserService := user.NewService(userRepo)
	userHandler := user.NewHandler(UserService)


	r := router.New(orgHandler, userHandler)

	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
