package main

import (
	"encoding/json"
	"net/http"    
	"strings"  
)

// GET /api/logout - Logout
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
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body BodyLoginAPILoginPost
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeJSON(w, 422, HTTPValidationError{
			Detail: []ValidationError{
				{
					Loc:  []any{"body"},
					Msg:  "Invalid request body",
					Type: "value_error",
				},
			},
		})
		return
	}


	if strings.TrimSpace(body.Username) == "" ||
		strings.TrimSpace(body.Password) == "" {

		writeJSON(w, 422, HTTPValidationError{
			Detail: []ValidationError{
				{
					Loc:  []any{"body", "username"},
					Msg:  "Field required",
					Type: "value_error.missing",
				},
				{
					Loc:  []any{"body", "password"},
					Msg:  "Field required",
					Type: "value_error.missing",
				},
			},
		})
		return
	}

	// Fake login success (DB ikke klar endnu)
	setUserID(w, "1")

	status := 200
	message := "logged in"

	writeJSON(w, 200, AuthResponse{
		StatusCode: &status,
		Message:    &message,
	})
}