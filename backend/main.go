package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Backend is healthy 🚀")
	}).Methods("GET")

	// User routes
	r.HandleFunc("/api/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/api/users", handlers.CreateUser).Methods("POST")

	// Start server
	port := "8080"
	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
