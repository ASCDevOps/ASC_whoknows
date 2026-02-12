package main

import (
	"encoding/json" // Needed for endpoints
	"net/http"      // http-pakke in go
)

type rootHandler struct{}

func (h *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Serve Root Page
	testTemplate.Execute(w, nil)
}

// GET /login - Serve Login Page
type loginHandler struct{}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Method not allowed",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":      true,
		"message": "Login endpoint ready. Use POST /api/login to authenticate.",
	})
}

// Helpers

func getUserID(r *http.Request) string {
	c, err := r.Cookie("user_id")
	if err != nil {
		return ""
	}
	return c.Value
}

func setUserID(w http.ResponseWriter, userID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",
		Value:    userID,
		Path:     "/",
		HttpOnly: true,
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}