# 16-04-2026

* I optimized the pipeline workflow so that:

```
PR → CI
merge → CI (master)
CI success → GHCR
GHCR success → deploy
deploy success → smoke test
```

* This was needed to stop double deployment and stops unneeded runner computing

# 23-04-2026

* Postman monitoring team free trial has ran out, so we have to find another monitoring system.
* This clashes well with prometheus & grafana, as we want to monitor the product.
* I have setup prometheus alertmanager with discord webhooks that sends a CRITICAL message, if the container has not been seen for 1 minute.
* Prometheus doesnt support environment from compose files as grafana does, therefore the prometheus file on github is a template.
