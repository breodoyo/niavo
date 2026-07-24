package main

import (
	"log"
	"net/http"

	"github.com/breodoyo/niavo/backend/internal/config"
	"github.com/breodoyo/niavo/backend/internal/router"
	
	
)

func main() {
	cfg := config.Load()

	r := router.New()

	log.Printf("Niavo API listening on %s", cfg.Port)

	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}