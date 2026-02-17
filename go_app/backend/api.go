package main

import (
	"encoding/json"
	"net/http"    
	"strings"
	"database/sql"  
)

// GET /api/logout - Logout
type logoutHandler struct{}

func (h *logoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := 200
	message := "logged out"

	resp := AuthResponse{
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
		writeJSON(w, 422, StandardResponse{
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

	writeJSON(w, 200, SearchResponse{Data: data})
}

// POST /api/login
type apiLoginHandler struct{}

func (h *apiLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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


	if strings.TrimSpace(body.Username) == "" ||
		strings.TrimSpace(body.Password) == "" {

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

	// Fake login success (DB ikke klar endnu)
	setUserID(w, "1")

	status := 200
	message := "logged in"

	writeJSON(w, 200, AuthResponse{
		StatusCode: &status,
		Message:    &message,
	})
}