package handlers

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
)

var searchTemplate = template.Must(template.ParseFiles("templates/layout.html", "templates/search.html"))

// HTML handler
type SearchHandler struct {
	DB *sql.DB
}

type SearchPageData struct {
	Query         string
	SearchResults []SearchResult
}

type SearchResult struct {
	Title       string
	URL         string
	Description string
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	data := SearchPageData{Query: query}

	if query != "" {
		data.SearchResults = queryPages(h.DB, query)
	}

	_ = searchTemplate.ExecuteTemplate(w, "layout", data)
}

// JSON API handler
type SearchAPIHandler struct {
	DB *sql.DB
}

type SearchResponse struct {
	Data []map[string]any `json:"data"`
}

func (h *SearchAPIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	results := queryPages(h.DB, query)

	data := make([]map[string]any, len(results))
	for i, r := range results {
		data[i] = map[string]any{
			"title":       r.Title,
			"url":         r.URL,
			"description": r.Description,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SearchResponse{Data: data})
}

// Shared query logic
func queryPages(db *sql.DB, query string) []SearchResult {
	rows, err := db.Query(`
		SELECT title, url, content
		FROM pages
		WHERE content ILIKE $1 OR title ILIKE $1
		ORDER BY title
		LIMIT 20
	`, "%"+query+"%")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var results []SearchResult
	for rows.Next() {
		var sr SearchResult
		var content string
		if err := rows.Scan(&sr.Title, &sr.URL, &content); err != nil {
			continue
		}
		if len(content) > 200 {
			sr.Description = content[:200] + "..."
		} else {
			sr.Description = content
		}
		results = append(results, sr)
	}
	return results
}
