# 09-02-2026

## Github
* Refactored documentation, so we have better structure going forward.
* Added a PR template

# 12-02-2026

## Github
### We have thought about splitting up the main.go into several files, primarily to avoid git-problems.
### Arguments for:
1. Less merge conflicts
2. More manageable file structure

### More Questions:
1. How many files should it be split into?
- When does it "stop"?

### Arguments against: 
1. The project will be more complicated
2. Duplication of code


### Final choice
1. We are deciding to make 3 .go files
2. **api.go** which serves api-endpoints, **http.go** which serves http and **main.go** which gathers and runs the server
