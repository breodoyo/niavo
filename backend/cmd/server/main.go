package main

import (
	"log"
	"net/http"

	"github.com/breodoyo/niavo/backend/internal/router"
	
)

func main() {
	r := router.New()

	log.Println("Niavo API listening on :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}