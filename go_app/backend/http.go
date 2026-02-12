package main

import ( 
	"net/http"      
	"html/template" 
)

type rootHandler struct{}
var testTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (h *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Serve Root Page
	testTemplate.Execute(w, nil)
}

// GET /login - Serve Login Page
type loginHandler struct{}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "templates/test.html")
}
