package handlers

import (
	"html/template"
	"net/http"
)

type RootHandler struct{}

var rootTemplate = template.Must(template.ParseFiles(
	"templates/layout.html",
	"templates/search.html",
))

func (*RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("session")
	if err == nil {
		_ = rootTemplate.ExecuteTemplate(w, "layout", map[string]any{
			"User": cookie.Value,
		})
	} else {
		_ = rootTemplate.ExecuteTemplate(w, "layout", nil)
	}
}
