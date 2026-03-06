package handlers

import (
	"database/sql"
	"html/template"
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

	if err := loginTemplate.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
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

	if err := r.ParseForm(); err != nil {
		writeJSON(w, 422, structs.HTTPValidationError{
			Detail: []structs.ValidationError{
				{
					Loc:  []any{"body"},
					Msg:  "Invalid request body",
					Type: "value_error",
				},
			},
		})
		return
	}

	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))

	if username == "" || password == "" {
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

		writeJSON(w, 500, structs.AuthResponse{
			StatusCode: &status,
			Message:    &msg,
		})
		return
	}

	var dbPassword string
	err := h.DB.QueryRow(
		"SELECT password FROM users WHERE username = ?",
		username,
	).Scan(&dbPassword)

	if err == sql.ErrNoRows {
		status := 401
		msg := "invalid credentials"

		writeJSON(w, 401, structs.AuthResponse{
			StatusCode: &status,
			Message:    &msg,
		})
		return
	}

	if err != nil {
		status := 500
		msg := "database error"

		writeJSON(w, 500, structs.AuthResponse{
			StatusCode: &status,
			Message:    &msg,
		})
		return
	}

	if dbPassword != password {
		status := 401
		msg := "invalid credentials"

		writeJSON(w, 401, structs.AuthResponse{
			StatusCode: &status,
			Message:    &msg,
		})
		return
	}

	status := 200
	msg := "logged in"

	writeJSON(w, 200, structs.AuthResponse{
		StatusCode: &status,
		Message:    &msg,
	})
}