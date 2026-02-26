package handlers

import (
	"html/template"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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
// ----- Types (kun nødvendige for at login.go kan stå alene) ----- //

type BodyLoginAPILoginPost struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	StatusCode *int    `json:"statusCode,omitempty"`
	Message    *string `json:"message,omitempty"`
}

type ValidationError struct {
	Loc  []any  `json:"loc"`
	Msg  string `json:"msg"`
	Type string `json:"type"`
}

type HTTPValidationError struct {
	Detail []ValidationError `json:"detail"`

}

// POST /api/login
type APILoginHandler struct{
	DB*sql.DB
}

func (h*APILoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	if strings.TrimSpace(body.Username) == "" || strings.TrimSpace(body.Password) == "" {
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

	if h.DB == nil {
		status := 500
		msg := "database not configured"
		writeJSON(w, 500, AuthResponse{
			StatusCode: &status,
			Message:    &msg,
		})
		return
	}

	// Slå bruger op
	var userID int
	var dbPassword string

	err := h.DB.QueryRow(
		"SELECT id, password FROM users WHERE username = ?",
		body.Username,
	).Scan(&userID, &dbPassword)

	if err == sql.ErrNoRows{
		status := 401
		msg :="invalid credentials"
		writeJSON(w, 401, AuthResponse{
			StatusCode: &status,
			Message: &msg,
		})
		return
	}

	if err != nil {
		status := 500
		msg := "database error"
		writeJSON(w, 500, AuthResponse{
			StatusCode: &status,
			Message:    &msg,
		})
		return
	}


	if dbPassword !=body.Password{
		status :=401
		msg :="invalid credentials"
		writeJSON(w, 401, AuthResponse{
			StatusCode: &status,
			Message:  &msg,
		})
		return
	}

	// Success
	setUserID(w, strconv.Itoa(userID))

	status := 200
	msg := "logged in"
	writeJSON(w, 200, AuthResponse{
		StatusCode: &status,
		Message:    &msg,
	})
}

// Helpers til POST /api/login

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