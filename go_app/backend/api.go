package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/http"
	"strings"
	"whoknows_backend/metrics"
	"whoknows_backend/structs"
)

// GET /api/search
type apiSearchHandler struct {
	DB *sql.DB
}

func (h *apiSearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	metrics.SearchRequestsTotal.Inc()

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

	normalized := strings.ToLower(strings.TrimSpace(q))
	if len(normalized) > 50 {
		normalized = normalized[:50]
	}

	metrics.SearchQueriesTotal.WithLabelValues(normalized).Inc()

	if h.DB == nil {
		metrics.SearchErrorsTotal.Inc()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rows, err := h.DB.Query(
		`SELECT title, url, language, last_updated, content
		FROM pages
		WHERE to_tsvector('english', title || ' ' || content) @@ plainto_tsquery($1)
		AND language = $2`,
		q, lang,
	)
	if err != nil {
		metrics.SearchErrorsTotal.Inc()
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

	metrics.SearchResultsTotal.Add(float64(len(data)))
	if len(data) == 0 {
		metrics.SearchNoResultsTotal.Inc()
	}

	writeJSON(w, 200, structs.SearchResponse{Data: data})
}

// Hjælper
func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
