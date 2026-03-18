package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	_ "modernc.org/sqlite"

	"whoknows_backend/security"
	"whoknows_backend/structs"
)

// GET /api/logout - Logout
type logoutHandler struct{}

func (*logoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	status := 200
	message := "logged out"

	resp := structs.AuthResponse{
		StatusCode: &status,
		Message:    &message,
	}

	writeJSON(w, http.StatusOK, resp)
}

// GET /api/search
type apiSearchHandler struct {
	DB *sql.DB
}

func (h *apiSearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query().Get("q")
	lang := r.URL.Query().Get("language")
	if lang == "" {
		lang = "en"
	}

	if strings.TrimSpace(q) == "" {
		status := 422
		msg := "q is required"
		writeJSON(w, http.StatusUnprocessableEntity, structs.StandardResponse{
			StatusCode: &status,
			Message:    &msg,
		})
		return
	}

	if h.DB == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rows, err := h.DB.Query(
		`SELECT title, url, language, last_updated, content
		 FROM pages
		 WHERE language = ? AND content LIKE ?`,
		lang, "%"+q+"%",
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = rows.Close()
	}()

	data := make([]map[string]any, 0)
	for rows.Next() {
		var title, url, language, content string
		var lastUpdated any

		if err := rows.Scan(&title, &url, &language, &lastUpdated, &content); err != nil {
			continue
		}

		data = append(data, map[string]any{
			"title":        title,
			"url":          url,
			"language":     language,
			"last_updated": lastUpdated,
			"content":      content,
		})
	}

	if err := rows.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, structs.SearchResponse{Data: data})
}

// POST /api/register
type registerHandlerAPI struct {
	db *sql.DB
}

func (h *registerHandlerAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]interface{}{
			"detail": []map[string]interface{}{
				{
					"loc":  []string{"body"},
					"msg":  "Invalid form data",
					"type": "value_error",
				},
			},
		})
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	password2 := r.FormValue("password2")

	if validationErr := validateRegisterInput(username, email, password, password2); validationErr != nil {
		writeJSON(w, http.StatusUnprocessableEntity, validationErr)
		return
	}

	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status_code": http.StatusInternalServerError,
			"message":     "Could not hash password",
		})
		return
	}

	stmt, err := h.db.Prepare(`INSERT INTO users (username, email, password) VALUES (?, ?, ?)`)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status_code": http.StatusInternalServerError,
			"message":     "Database error",
		})
		return
	}
	defer func() {
		_ = stmt.Close()
	}()

	_, err = stmt.Exec(username, email, hashedPassword)
	if err != nil {
		writeJSON(w, http.StatusConflict, map[string]interface{}{
			"status_code": http.StatusConflict,
			"message":     "Username or email already exists",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"status_code": http.StatusOK,
		"message":     "User registered successfully",
	})
}

func validateRegisterInput(username, email, password, password2 string) map[string]interface{} {
	if username == "" {
		return map[string]interface{}{
			"detail": []map[string]interface{}{
				{
					"loc":  []string{"username"},
					"msg":  "Field required",
					"type": "value_error.missing",
				},
			},
		}
	}

	if email == "" {
		return map[string]interface{}{
			"detail": []map[string]interface{}{
				{
					"loc":  []string{"email"},
					"msg":  "Field required",
					"type": "value_error.missing",
				},
			},
		}
	}

	if password == "" {
		return map[string]interface{}{
			"detail": []map[string]interface{}{
				{
					"loc":  []string{"password"},
					"msg":  "Field required",
					"type": "value_error.missing",
				},
			},
		}
	}

	if password2 != "" && password != password2 {
		return map[string]interface{}{
			"detail": []map[string]interface{}{
				{
					"loc":  []string{"password2"},
					"msg":  "Passwords do not match",
					"type": "value_error",
				},
			},
		}
	}

	return nil
}

// Hjælper
func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
	}
}
