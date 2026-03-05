package handlers

import (
	"html/template"
	"net/http"
	"database/sql"
	"strings"
)

type RegisterHandler struct{
	DB *sql.DB
}

var registerTemplate = template.Must(template.ParseFiles("templates/test.html"))

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return	
	}

	username := strings.TrimSpace(r.FormValue("username"))
	email := strings.TrimSpace(r.FormValue("email"))
	password := strings.TrimSpace(r.FormValue("password"))
	password1 := strings.TrimSpace(r.FormValue("password1"))

	if username == "" || email == "" || password == "" {
		http.Error(w, "missing fields", http.StatusBadRequest)
		return
	}

	if password != password1 {
		http.Error(w, "Password do not match", http.StatusBadRequest)
		return
	}

	_, err = h.DB.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		username, email, password,
	)

	if err != nil{
		http.Error(w, "User creation failed!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"statusCode":200,"message":"User registered"}`))
}


