# 14-04-2026

* Added continous_deployment.yml
* It's not Full CD, because nginx and certbot is setup manually on the server. Therefore it wont work on a completely new server without setting those two up.
* Also doesn't log in to GHCR, that is already setup on the VM.

* Added smoke_testing.yml
* Smoke testing has also been setup, after Deployment is run.
* To check if website is still alive after deploying.

# 10-05-2026
* Since monitoring is already part of our DevOps workflow, we chose to measure performance improvements
   through Prometheus and Grafana rather than    external benchmarking tools.
* We implemented request duration metrics for our /api/search endpoint and used Grafana to compare the average response
  times before and after adding a PostgreSQL GIN index for full text search.
  This allowed us to monitor the impact of the index directly from the API and user perspective while staying within our existing observability setup.
