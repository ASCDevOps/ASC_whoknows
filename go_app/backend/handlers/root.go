package handlers

import (
	"html/template"
	"net/http"
)

type RootHandler struct{}

var rootTemplate = template.Must(template.ParseFiles(
	"templates/layout.html",
	"templates/about.html",
	"templates/login.html",
	"templates/register.html",
	"templates/search.html",

))

func (*RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Just render the page (no data needed anymore)
	err := rootTemplate.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
