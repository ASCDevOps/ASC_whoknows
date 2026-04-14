# 14-04-2026

* Added continous_deployment.yml
* It's not Full CD, because nginx and certbot is setup manually on the server. Therefore it wont work on a completely new server without setting those two up.
* Also doesn't log in to GHCR, that is already setup on the VM.

* Added smoke_testing.yml
* Smoke testing has also been setup, after Deployment is run.
* To check if website is still alive after deploying.
