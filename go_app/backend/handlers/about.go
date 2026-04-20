package handlers

import (
	"html/template"
	"net/http"
)

type AboutHandler struct{}

var aboutTemplate = template.Must(template.ParseFiles("templates/layout.html", "templates/about.html"))

func (*AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("session")
	if err == nil {
		_ = aboutTemplate.ExecuteTemplate(w, "layout", map[string]any{
			"User": cookie.Value,
		})
	} else {
		_ = aboutTemplate.ExecuteTemplate(w, "layout", nil)
	}
}
