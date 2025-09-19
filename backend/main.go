package main

import (
	"log"
	"net/http"

	"github.com/Fayez73/go-app-jenkins/backend/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// API route
	r.HandleFunc("/api/test", handlers.TestHandler).Methods("GET")

	log.Println("Starting backend on :4000")
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}
}
