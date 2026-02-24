package main

import (
	"html/template"
	"net/http"
)

type rootHandler struct{}

var rootTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (*rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_ = rootTemplate.Execute(w, nil)
}

type registerHandler struct{}

var registerTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (*registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_ = registerTemplate.Execute(w, nil)
}

type loginHandler struct{}

var loginTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (*loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_ = loginTemplate.Execute(w, nil)
}
