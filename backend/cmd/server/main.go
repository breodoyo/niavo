package main

import (
	"log"
	"net/http"

	"github.com/breodoyo/niavo/backend/internal/config"
	"github.com/breodoyo/niavo/backend/internal/database"
	"github.com/breodoyo/niavo/backend/internal/organization"
	"github.com/breodoyo/niavo/backend/internal/router"
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
	repo := organization.NewRepository(db)
	service := organization.NewService(repo)
	handler := organization.NewHandler(service)

	r := router.New(handler)

	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
