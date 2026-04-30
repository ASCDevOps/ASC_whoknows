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

# 30-04-2026
* We have chosen to focus our monitoring on a few key areas that together provide a clear picture of how the application is used and how well it performs in practice.
* First, we monitor the search functionality as a whole. This includes how often users perform searches, what they search for, and how successful those searches are.
  By grouping these metrics, we gain insight into user behavior and can evaluate whether the search feature is meeting user needs. For example, frequent searches combined with many unsuccessful results may indicate missing content or a need to improve the search logic.
* Secondly, we monitor overall user activity. By tracking how users interact with different parts of the application, we can better understand usage patterns and identify which features are most important.
  This helps us prioritize improvements and ensure that development efforts are focused on what provides the most value to users.
* We also monitor our API endpoints to understand how the system is being used from a technical perspective.
  This includes tracking request counts and identifying which endpoints are most frequently accessed. These metrics help us detect unusual traffic patterns, measure system load, and ensure that our backend services are performing as expected.
* In addition, we focus on the quality and reliability of the system by tracking errors and failed operations.
  This allows us to quickly identify technical issues in the backend, database, or request handling, and respond before they impact the user experience significantly.
* We have deliberately chosen not to include direct monitoring of server uptime and downtime within the dashboard.
  Instead, we rely on alerting through Alertmanager, which sends notifications to Discord whenever issues occur. Since we are actively using Discord, this approach allows us to receive immediate alerts and respond quickly to incidents without needing to constantly monitor uptime metrics manually.
* Overall, this approach provides a balanced view of both user behavior and system performance.
  It supports a DevOps mindset, where we continuously observe how the system is used in real-world scenarios and use that insight to improve both functionality and reliability over time.
