package handlers

import (
	"html/template"
	"net/http"
)

type RegisterHandler struct{}

var registerTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (*RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_ = registerTemplate.Execute(w, nil)
}
