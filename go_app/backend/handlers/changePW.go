package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"whoknows_backend/security"
	"whoknows_backend/structs"
)

type ChangePasswordHandler struct {
	DB *sql.DB
}

func (h *ChangePasswordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body structs.BodyChangePassword
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(body.Username) == "" ||
		strings.TrimSpace(body.NewPassword) == "" ||
		strings.TrimSpace(body.ConfirmPassword) == "" {
		http.Error(w, "missing fields", http.StatusBadRequest)
		return
	}

	if body.NewPassword != body.ConfirmPassword {
		http.Error(w, "passwords do not match", http.StatusBadRequest)
		return
	}

	hashedPassword, err := security.HashPassword(body.NewPassword)
	if err != nil {
		http.Error(w, "failed to hash password", http.StatusInternalServerError)
		return
	}

	_, err = h.DB.Exec(
		"UPDATE users SET password = ?, must_change_password = 0 WHERE username = ?",
		hashedPassword,
		body.Username,
	)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"message":"password changed successfully"}`))
}
