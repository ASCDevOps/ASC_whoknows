package handlers

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	"whoknows_backend/security"
	"whoknows_backend/structs"
)

type LoginHandler struct{}

var loginTemplate = template.Must(template.ParseFiles("templates/layout.html", "templates/login.html"))

func (*LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_ = loginTemplate.ExecuteTemplate(w, "layout", nil)
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
	var mustChangePassword int
	err := h.DB.QueryRow(
		"SELECT password, must_change_password FROM users WHERE username = ?",
		body.Username,
	).Scan(&dbPassword, &mustChangePassword)

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

	// Password check (HashingPassword)
	if err := security.CheckPassword(body.Password, dbPassword); err != nil {
		// Fallback for gamle plaintext-passwords
		if dbPassword == body.Password {
			hashedPassword, hashErr := security.HashPassword(body.Password)
			if hashErr != nil {
				status := 500
				msg := "failed to upgrade password"
				writeJSON(w, 500, structs.AuthResponse{
					StatusCode: &status,
					Message:    &msg,
				})
				return
			}

			_, updateErr := h.DB.Exec(
				"UPDATE users SET password = ?, must_change_password = 1 WHERE username = ?",
				hashedPassword,
				body.Username,
			)
			if updateErr != nil {
				status := 500
				msg := "failed to upgrade password"
				writeJSON(w, 500, structs.AuthResponse{
					StatusCode: &status,
					Message:    &msg,
				})
				return
			}

			mustChangePassword = 1
		} else {
			status := 401
			msg := "invalid credentials"
			writeJSON(w, 401, structs.AuthResponse{
				StatusCode: &status,
				Message:    &msg,
			})
			return
		}
	}
	status := 200
	msg := "logged in"
	requiresPasswordChange := false

	if mustChangePassword == 1 {
		msg = "password change required"
		requiresPasswordChange = true
	}

	writeJSON(w, 200, structs.AuthResponse{
		StatusCode:             &status,
		Message:                &msg,
		RequiresPasswordChange: &requiresPasswordChange,
	})

}

// Helpers til POST /api/login
func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
