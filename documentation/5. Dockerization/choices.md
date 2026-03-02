# 02-03-2026

## Dockerization
* .dockerignore created by configs.sh
* The Dockerfile currently doesn't have a set WORKDIR because it messes with paths, specifically html.
* I have chosen to have chaos in the file structure inside the container in order to meet the deadline.
* The Dockerfile should be rewritten with a specific WORKDIR, so the file structure is more clean.