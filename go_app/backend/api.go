package main

import (
	"encoding/json" // Needed for endpoints
	"net/http"      // http-pakke in go
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
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Method not allowed",
		})
		return
	}

	if err := r.ParseForm(); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Bad request",
		})
		return
	}
	// "Session" via cookie
	setUserID(w, "1")

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":     true,
		"userId": "1",
	})
}

type apiMeHandler struct{}

func (h *apiMeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Method not allowed",
		})
		return
	}

	userID := getUserID(r)
	if userID == "" {
		writeJSON(w, http.StatusUnauthorized, map[string]any{
			"ok":     false,
			"authed": false,
			"error":  "Not logged in",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":     true,
		"authed": true,
		"userId": userID,
	})
}
