package handlers

import (
	"html/template"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"whoknows_backend/structs"
)

type LoginHandler struct{}

var loginTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (*LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_ = loginTemplate.Execute(w, nil)
}

// POST /api/login
type APILoginHandler struct {
	DB *sql.DB
}

func (h *APILoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body structs.BodyLoginAPILoginPost
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeJSON(w, 422, structs.HTTPValidationError{
			Detail: []structs.ValidationError{
				{Loc: []any{"body"}, Msg: "Invalid request body", Type: "value_error"},
			},
		})
		return
	}

	if strings.TrimSpace(body.Username) == "" || strings.TrimSpace(body.Password) == "" {
		writeJSON(w, 422, structs.HTTPValidationError{
			Detail: []structs.ValidationError{
				{Loc: []any{"body", "username"}, Msg: "Field required", Type: "value_error.missing"},
				{Loc: []any{"body", "password"}, Msg: "Field required", Type: "value_error.missing"},
			},
		})
		return
	}

	if h.DB == nil {
		status := 500
		msg := "database not configured"
		writeJSON(w, 500, structs.AuthResponse{StatusCode: &status, Message: &msg})
		return
	}

	// Slå bruger op (kun password er nødvendigt for login)
	var dbPassword string
	err := h.DB.QueryRow(
		"SELECT password FROM users WHERE username = ?",
		body.Username,
	).Scan(&dbPassword)

	if err == sql.ErrNoRows {
		status := 401
		msg := "invalid credentials"
		writeJSON(w, 401, structs.AuthResponse{StatusCode: &status, Message: &msg})
		return
	}
	if err != nil {
		status := 500
		msg := "database error"
		writeJSON(w, 500, structs.AuthResponse{StatusCode: &status, Message: &msg})
		return
	}

	// Password check (plaintext lige nu)
	if dbPassword != body.Password {
		status := 401
		msg := "invalid credentials"
		writeJSON(w, 401, structs.AuthResponse{StatusCode: &status, Message: &msg})
		return
	}

	// Success
	status := 200
	msg := "logged in"
	writeJSON(w, 200, structs.AuthResponse{StatusCode: &status, Message: &msg})
}
// Helpers til POST /api/login

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}