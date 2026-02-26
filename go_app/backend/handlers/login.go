package handlers

import (
	"html/template"
	"net/http"
)

type LoginHandler struct{}

var loginTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (*LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_ = loginTemplate.Execute(w, nil)
}