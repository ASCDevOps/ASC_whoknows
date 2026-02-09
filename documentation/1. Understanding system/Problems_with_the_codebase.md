# 05-02-2026

## Notes
* Environment variables doesn't seem to be setup correctly in html templates, neither css link.

* I think this is because the dependencies in requirements.txt are discontinued and isn't supported by Python3 which app.py has been converted to via 2to3.

* Also in app.py a problem occurs with flask "import "flask" could not be resolved"

* Doesn't look like regular html, have to look into what is used. (flask?)

## Styling
* The issues that may arise typically stem from not respecting the existing legacy code and its associated structure and styling.

* If too many changes are made, it can lead to more problems than necessary, both functionally and visually. 

* Apart from that, there are no major errors, but mainly minor bugs that may occur due to the way the legacy code is structured.

## Critical Problems
1. SQL injection
    - build with string
2. Unsafe password hashing 
    - not compadebel with moden systems
3. Hardcoded secret key
    - the application uses a fixed secret key for session management.
4. Application can terminate abruptly 
    - If the database is missing, the server terminates instead of handling the error gracefully
5. Error handling is exposed directly
    - Potential errors may result in unhandled exceptions and exposed stack traces.
6. Missing input validation
    - User input is used directly without sufficient validation or sanitization.