package handlers

import (
	"html/template"
	"net/http"
	"database/sql"
	"strings"

	"whoknows_backend/structs"
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

	if req.Username == "" || req.Email == "" || req.Password == "" {
		msg := "Missing fields"
		res := structs.AuthResponse{
			StatusCode: intPtr(422),
			Message:	&msg,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(res)
		return
	}

	if req.Password != req.Password1{
		msg := "Passwords do not match!"
		res := structs.AuthResponse{
			StatusCode: intPtr(422),
			Message:	&msg, 
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(res)
		return
	} 

	_, err = h.DB.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		req.Username, req.Email, req.Password,
	)

	if err != nil{
		msg := "User creation failed!"
		res := structs.AuthResponse{
			statusCode: intPtr(500),
			Message:	&msg,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return

	msg := "User Registered"
	res := structs.AuthResponse{
		StatusCode: intPtr(200),
		Message:	&msg,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

	}

func intPtr(i int) *int {
	return &i
}


