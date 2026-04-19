package main

import (
	"database/sql"
	"encoding/json"
	_ "modernc.org/sqlite"
	"net/http"
	"strings"
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
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

	// q er required -> 422 hvis den mangler
	if strings.TrimSpace(q) == "" {
		status := 422
		msg := "q is required"
		writeJSON(w, 422, structs.StandardResponse{
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
	`SELECT p.title, p.url, p.language, p.last_updated, p.content
	 FROM pages_fts
	 JOIN pages p ON p.rowid = pages_fts.rowid
	 WHERE pages_fts MATCH ?
	   AND p.language = ?`,
	q, lang,
)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	data := make([]map[string]any, 0)
	for rows.Next() {
		var title, url, language, content string
		var lastUpdated any // kan være null
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

	writeJSON(w, 200, structs.SearchResponse{Data: data})
}

// POST /api/register
type registerHandlerAPI struct {
	db *sql.DB
}

func (h *registerHandlerAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Checks if method is POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parses form,
	if err := r.ParseForm(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]interface{}{
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

	// Gets form-data
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	password2 := r.FormValue("password2")

	// Validates form-data, checks for required, checks if passwords are matching
	if username == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"detail": []map[string]interface{}{
				{
					"loc":  []string{"username"},
					"msg":  "Field required",
					"type": "value_error.missing",
				},
			},
		})
		return
	}
	if email == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"detail": []map[string]interface{}{
				{
					"loc":  []string{"email"},
					"msg":  "Field required",
					"type": "value_error.missing",
				},
			},
		})
		return
	}
	if password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"detail": []map[string]interface{}{
				{
					"loc":  []string{"password"},
					"msg":  "Field required",
					"type": "value_error.missing",
				},
			},
		})
		return
	}
	if password2 != "" && password != password2 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"detail": []map[string]interface{}{
				{
					"loc":  []string{"password2"},
					"msg":  "Passwords do not match",
					"type": "value_error",
				},
			},
		})
		return
	}

	stmt, err := h.db.Prepare(`INSERT INTO users (username, email, password) VALUES (?, ?, ?)`)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusInternalServerError,
			"message":     "Failed to prepare database statement",
		})
		return
	}
	defer stmt.Close()

	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusInternalServerError,
			"message":     "Failed to hash password",
		})
		return
	}

	// error handling for UNIQUE (username & email) in SQL
	_, err = stmt.Exec(username, email, hashedPassword)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusConflict,
			"message":     "Username or email already exists",
		})
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status_code": http.StatusOK,
		"message":     "User registered successfully",
	})
}

// Hjælper
func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
