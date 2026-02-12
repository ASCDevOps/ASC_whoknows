	package main

	import (
		"encoding/json"
		"net/http"
	)

	// GET /login - Serve Login Page
		 //mux.Handle("/login", &loginHandler{}) // <-- NY

		// POST /api/login - Login
		 //mux.Handle("/api/login", &apiLoginHandler{}) // <-- NY


	// Skal ligge under min logout handler


// GET /login - Serve Login Page
type loginHandler struct{}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Method not allowed",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":      true,
		"message": "Login endpoint ready. Use POST /api/login to authenticate.",
	})
}

	// POST /api/login
type apiLoginHandler struct{}

func (h *apiLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Method not allowed",
		})
		return
	}

	if err := r.ParseForm(); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Bad request",
		})
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Midlertidig test-login (kobles på DB senere)
	if username != "chris" {
		writeJSON(w, http.StatusUnauthorized, map[string]any{
			"ok":    false,
			"error": "Invalid username",
		})
		return
	}
	if password != "secret" {
		writeJSON(w, http.StatusUnauthorized, map[string]any{
			"ok":    false,
			"error": "Invalid password",
		})
		return
	}

	// "Session" via cookie
	setUserID(w, "1")

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":     true,
		"userId": "1",
	})
}

// GET /api/me (optional) - check if logged in
type apiMeHandler struct{}

func (h *apiMeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Method not allowed",
		})
		return
	}

	userID := getUserID(r)
	if userID == "" {
		writeJSON(w, http.StatusUnauthorized, map[string]any{
			"ok":      false,
			"authed":  false,
			"error":   "Not logged in",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":     true,
		"authed": true,
		"userId": userID,
	})
}

// Helpers

func getUserID(r *http.Request) string {
	c, err := r.Cookie("user_id")
	if err != nil {
		return ""
	}
	return c.Value
}

func setUserID(w http.ResponseWriter, userID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",
		Value:    userID,
		Path:     "/",
		HttpOnly: true,
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}