package handlers

import (
	"encoding/json"
	"net/http"
)

type TestResponse struct {
	Message string `json:"message"`
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TestResponse{
		Message: "Backend is working!",
	})
}
