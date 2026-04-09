package handlers

import (
	"html/template"
	"net/http"
)

type RootHandler struct{}

var rootTemplate = template.Must(template.ParseFiles("templates/layout.html"))

func (*RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := rootTemplate.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
