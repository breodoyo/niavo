package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", homeHandler)
	log.Println("Niavo API listening on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
	
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string {
		"name": "Niavo API",
		"version": "0.1.0",
		"status": "running",
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode respone", http.StatusInternalServerError)
	}
}