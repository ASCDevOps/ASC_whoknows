package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"database/sql"
	
	"whoknows_backend/structs"
)

type RegisterHandler struct {
	DB *sql.DB
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	req := structs.BodyRegisterAPIRegisterPost{
		Username:  strings.TrimSpace(r.FormValue("username")),
		Email:     strings.TrimSpace(r.FormValue("email")),
		Password:  strings.TrimSpace(r.FormValue("password")),
		Password2: strings.TrimSpace(r.FormValue("password2")),
	}

	// Validation
	if req.Username == "" || req.Email == "" || req.Password == "" {
		msg := "missing fields"

		writeJSON(w, http.StatusUnprocessableEntity, structs.AuthResponse{
			StatusCode: intPtr(422),
			Message:    &msg,
		})
		return
	}

	if req.Password != req.Password2 {
		msg := "passwords do not match"

		writeJSON(w, http.StatusUnprocessableEntity, structs.AuthResponse{
			StatusCode: intPtr(422),
			Message:    &msg,
		})
		return
	}

	if h.DB == nil {
		msg := "database not configured"

		writeJSON(w, http.StatusInternalServerError, structs.AuthResponse{
			StatusCode: intPtr(500),
			Message:    &msg,
		})
		return
	}

	_, err := h.DB.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		req.Username, req.Email, req.Password,
	)

	if err != nil {
		msg := "user creation failed"

		writeJSON(w, http.StatusInternalServerError, structs.AuthResponse{
			StatusCode: intPtr(500),
			Message:    &msg,
		})
		return
	}

	msg := "user registered"

	writeJSON(w, http.StatusOK, structs.AuthResponse{
		StatusCode: intPtr(200),
		Message:    &msg,
	})
}

func intPtr(i int) *int {
	return &i
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "encoding error", http.StatusInternalServerError)
	}
}