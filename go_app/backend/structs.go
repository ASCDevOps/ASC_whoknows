package main

// json: matching names for json
// omiempty = can be nil

// Request Bodies
// SKRIV HVILKET ENDPOINT DET BRUGES TIL HER
type BodyLoginAPILoginPost struct {
}

// SKRIV HVILKET ENDPOINT DET BRUGES TIL HER
type BodyRegisterAPIRegisterPost struct {
}

// Responses
// GET /api/logout,
type AuthResponse struct {
	StatusCode *int    `json:"statusCode,omitempty"`
	Message    *string `json:"message,omitempty"`
}

// SKRIV HVILKET ENDPOINT DET BRUGES TIL HER
type StandardResponse struct {
}

// SKRIV HVILKET ENDPOINT DET BRUGES TIL HER
type SearchResponse struct {
}

// Validation Errors
// SKRIV HVILKET ENDPOINT DET BRUGES TIL HER
type ValidationError struct {
}

// SKRIV HVILKET ENDPOINT DET BRUGES TIL HER
type HTTPValidationError struct {
}

// SKRIV HVILKET ENDPOINT DET BRUGES TIL HER
type RequestValidationError struct {
}
