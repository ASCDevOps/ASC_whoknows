package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"whoknows_backend/structs"
)

type RegisterHandler struct {
	DB *sql.DB
}

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

	req := structs.BodyRegisterAPIRegisterPost{
		Username:  strings.TrimSpace(r.FormValue("username")),
		Email:     strings.TrimSpace(r.FormValue("email")),
		Password:  strings.TrimSpace(r.FormValue("password")),
		Password2: strings.TrimSpace(r.FormValue("password2")),
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		msg := "Missing fields"

		res := structs.AuthResponse{
			StatusCode: intPtr(422),
			Message:    &msg,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(res)
		return
	}

	if req.Password != req.Password2 {
		msg := "Passwords do not match!"

		res := structs.AuthResponse{
			StatusCode: intPtr(422),
			Message:    &msg,
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

	if err != nil {
		msg := "User creation failed!"

		res := structs.AuthResponse{
			StatusCode: intPtr(500),
			Message:    &msg,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	msg := "User registered"

	res := structs.AuthResponse{
		StatusCode: intPtr(200),
		Message:    &msg,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func intPtr(i int) *int {
	return &i
}