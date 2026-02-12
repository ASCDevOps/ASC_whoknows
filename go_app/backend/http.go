package main

import (
	"encoding/json" // Needed for endpoints
	"html/template" // templating-pakke in go
	"net/http"      // http-pakke in go
)

type rootHandler struct{}

var rootTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (h *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Serve Root Page
	rootTemplate.Execute(w, nil)
}

type registerHandler struct{}

var registerTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (h *registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Serve Register Page
	registerTemplate.Execute(w, nil)
}

// GET /login - Serve Login Page
type loginHandler struct{}

var loginTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_ = loginTemplate.Execute(w, nil)
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
