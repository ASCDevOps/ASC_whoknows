package handlers

import (
	"html/template"
	"net/http"
)

type RootHandler struct{}

var rootTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (*RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := rootTemplate.Execute(w, nil); err != nil {
		http.Error(w, "template render error", http.StatusInternalServerError)
		return
	}
}