# 02-03-2026

## Dockerization
* In order not to leak the database, I have chosen to use volumes and mount the database onto the container.
* .dockerignore created by configs.sh
* The Dockerfile currently doesn't have a set WORKDIR because it messes with paths, specifically html.
* I have chosen to have chaos in the file structure inside the container in order to meet the deadline.
* The Dockerfile should be rewritten with a specific WORKDIR, so the file structure is more clean.

## Postman Monitoring
* Added Postman Monitoring
* We are now monitoring frontend http every hour.
* Will receive email if failed
* We have no way of mocking the DB, and dont want to fill the DB up with dummy data from postman monitoring.