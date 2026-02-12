# 09-02-2026

## Github
* Refactored documentation, so we have better structure going forward.
* Added a PR template

# 11-02-2026

* During the migration process, I made several architectural decisions to simplify the transition. 
* Instead of rewriting all Jinja2 templates to Go’s template engine, I chose to treat Go as A pure API backend. This allowed me to keep the existing frontend structure while focusing on backend functionality. 
* I also decided to implement authentication using HTTP cookies rather  
* Than introducing external session libraries. Additionally, I structured the Go application with clear separation between routing and handler logic to align with Go’s explicit design 
* philosophy. These choices resulted in a cleaner, more maintainable backend architecture.