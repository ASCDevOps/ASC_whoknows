# 11-02 - 2026
* During the migration from a Flask application to Go, I faced several structural and architectural challenges. 
* First, I had to resolve environment issues, including installing Go and configuring the system PATH correctly. 
* I then had to adapt to Go’s stricter application structure, including the requirement for a single main() function and explicit route registration. 
* A major challenge was the incompatibility between Jinja2 templates and Go’s html/template engine, which forced a shift away from server-side rendering. 
* This required rethinking the login flow and moving toward a more API-driven backend design. Additionally, session handling had to be implemented manually in Go, unlike Flask’s built-in request context and session management.