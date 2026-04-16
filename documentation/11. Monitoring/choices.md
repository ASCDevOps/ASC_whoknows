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
