package main

import (
	"log"
	"net/http"

	"github.com/breodoyo/niavo/backend/internal/config"
	"github.com/breodoyo/niavo/backend/internal/router"
	
	
)

func main() {
	cfg := config.Load()

	log.Printf(
		"%s (%s) listening on %s", 
		cfg.AppName,
		cfg.Env,
		cfg.Port,
	)

	r := router.New()

	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}