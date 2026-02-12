package main

import (
	"encoding/json" 
	"net/http"      
)

// Schemas
// json: matching names for json
// omiempty = can be nil
type AuthResponse struct {
	StatusCode *int    `json:"statusCode,omitempty"`
	Message    *string `json:"message,omitempty"`
}


type logoutHandler struct{}

func (h *logoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := 200
	message := "logged out"

	resp := AuthResponse{
		StatusCode: &status,
		Message:    &message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// POST /api/login
type apiLoginHandler struct{}

func (h *apiLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Simpel fake validering
	if username == "admin" && password == "1234" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"statusCode":200,"message":"Logged in"}`))
		return
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}