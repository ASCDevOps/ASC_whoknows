package structs

// json: matching names for json
// omiempty = can be nil

// Request Bodies
// Post /api/login
type BodyLoginAPILoginPost struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

// GET /api/search til 442 response
type StandardResponse struct {
	StatusCode *int    `json:"statusCode,omitempty"`
	Message    *string `json:"message,omitempty"`
}

// GET /api/search til 200 response
type SearchResponse struct {
	Data []map[string]any `json:"data"`
}

// Validation Errors

// 422 validation Error - Post api/login
type ValidationError struct {
	Loc  []any  `json:"loc"` // (string | integer)
	Msg  string `json:"msg"`
	Type string `json:"type"`
}

// 422 validation Error wrapper . Post api/login
type HTTPValidationError struct {
	Detail []ValidationError `json:"detail"`
}

// SKRIV HVILKET ENDPOINT DET BRUGES TIL HER
type RequestValidationError struct {
}
