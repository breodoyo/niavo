package main

import (
	"log"
	"net/http"

	"github.com/breodoyo/niavo/backend/internal/app"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", app.HomeHandler)

	log.Println("Niavo API listening on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}