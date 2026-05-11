# WhoKnows

![Build](https://img.shields.io/github/actions/workflow/status/ASCDevOps/ASC_whoknows/continous_integration.yml?branch=master&label=Build&style=for-the-badge)
![Go Version](https://img.shields.io/github/go-mod/go-version/ASCDevOps/ASC_whoknows?filename=go_app/backend/go.mod&style=for-the-badge)
![Release](https://img.shields.io/github/v/release/ASCDevOps/ASC_whoknows?style=for-the-badge)
[![Go CI](https://github.com/ASCDevOps/ASC_whoknows/actions/workflows/continous_integration.yml/badge.svg)](https://github.com/ASCDevOps/ASC_whoknows/actions/workflows/continous_integration.yml)

WhoKnows is a full-stack web search engine. Users can register, log in, and search an indexed content database. The project emphasises production-grade infrastructure: containerised deployment, monitoring, and automated CI/CD.

---

## Project Structure (MonoRepo)

We are using a MonoRepo project structure with 4 major folders

1. /documentation
2. /go_app
3. /monitor
4. /python_app

---

## Architecture

WhoKnows runs across two VMs:

``` Architecture
                        Internet
                           │
                    ┌──────▼──────┐
                    │    nginx    │  (TLS, reverse proxy)
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │   Go App    │  :8080
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │ PostgreSQL  │
                    └─────────────┘

        App VM                        Monitoring VM
  ┌─────────────────┐           ┌───────────────────────┐
  │ app             │           │ Prometheus   :9090    │
  │ nginx           │◄──scrape──│ Grafana      :3000    │
  │ postgres        │           │ Alertmanager :9093    │
  │ node-exporter   │           │ alertmanager-discord  │
  │ cAdvisor  :8081 │           └───────────────────────┘
  └─────────────────┘
```

The app VM runs the application stack. The monitoring VM scrapes metrics from node-exporter (host metrics) and cAdvisor (container metrics) on the app VM, as well as the Go app's own `/metrics` endpoint.

---

## Tech Stack

### Application

- [Go](https://golang.org/) — HTTP server, handlers, business logic
- [PostgreSQL](https://www.postgresql.org/) — primary database (users, pages)
- [nginx](https://nginx.org/) — reverse proxy, TLS termination via Let's Encrypt

### Infrastructure

- [Docker & Docker Compose](https://docs.docker.com/compose/) — containerisation
- [GitHub Actions](https://github.com/features/actions) — CI/CD pipeline
- [GitHub Container Registry](https://ghcr.io) — image hosting (`ghcr.io/ascdevops/asc_whoknows`)

### Monitoring

- [Prometheus](https://prometheus.io/) — metrics collection and alerting rules
- [Grafana](https://grafana.com/) — dashboards
- [Alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager/) — alert routing
- [Discord](https://discord.com/) — alert notifications via webhook

### Testing

- [Playwright](https://playwright.dev/) — end-to-end tests

---

## Getting Started

### Prerequisites

- Docker and Docker Compose installed on the target VM
- A `.env` file in place (see [Configuration](#configuration--secrets))

### App

```bash
# Clone the repository
git clone https://github.com/ascdevops/asc_whoknows.git
cd asc_whoknows/go_app/backend

# Copy and fill in the env file
cp .env.example .env

# Development (builds locally)
docker compose -f docker-compose.dev.yml up --build

# Production (pulls from ghcr.io)
docker compose -f docker-compose.prod.yml up -d
```

The app will be available at `http://localhost:8080` in development.

### Monitor

The monitoring stack runs on a separate VM.

```bash
cd monitor

# Copy and fill in the env file
cp .env.example .env

# Start the stack
docker compose up -d
```

| Service      | Port  |
|--------------|-------|
| Prometheus   | 9090  |
| Grafana      | 3000  |
| Alertmanager | 9093  |

---

## Development Setup

This project uses pre-commit hooks to enforce basic code hygiene and Go formatting before each commit. Lightweight checks run locally, while full linting and testing are handled in CI.

### Initialize pre-commit

Run the following commands from the repository root.

**Windows (PowerShell)** — requires Python

```bash
pip install pre-commit
pre-commit install
```

**macOS** — using Homebrew

```bash
brew install pre-commit
pre-commit install
```

To update hooks:

```bash
pre-commit autoupdate
```

After installation, hooks will automatically run on every `git commit`.

---

## Configuration & Secrets

### App — `go_app/backend/.env`

| Variable            | Description                              |
|---------------------|------------------------------------------|
| `DATABASE_URL`      | PostgreSQL connection string             |
| `POSTGRES_PASSWORD` | Password for the `whoknows` DB user      |
| `ADMIN_USERNAME`    | Username for the seeded admin account    |
| `ADMIN_EMAIL`       | Email for the seeded admin account       |
| `ADMIN_PASSWORD`    | Password for the seeded admin account    |

### Monitor — `monitor/.env`

| Variable          | Description                                      |
|-------------------|--------------------------------------------------|
| `GRAFANA_USER`    | Grafana admin username                           |
| `GRAFANA_PASSWORD`| Grafana admin password                           |
| `DISCORD_WEBHOOK` | Discord webhook URL for alert notifications      |
| `APP_VM_IP`       | IP address of the app VM (used in scrape config) |

---

## CI/CD Pipeline

The pipeline is defined in `.github/workflows/` with the following trigger logic:

``` Pipeline
PR opened/updated  →  CI (build & test only)

Merge to master    →  CI → push to GHCR → deploy to app VM → smoke test
```

Each stage gates the next — a failed CI run will not produce an image, a failed push will not trigger a deploy, and a failed deploy will not run the smoke test.

---

## Monitoring & Alerts

Prometheus scrapes three targets on the app VM:

| Job                | Target          | What it covers                         |
|--------------------|-----------------|-----------------------------------     |
| `whoknows-backend` | `:8080/metrics` | App-level metrics                      |
| `cadvisor`         | `:8081`         | Docker container metrics               |
| `node-exporter`    | `:9100`         | Host metrics (CPU, RAM, disk, network) |

**Alert rules** (defined in `monitor/alertmanager/alerts.yml`):

| Alert              | Condition                      | Severity |
|--------------------|--------------------------------|----------|
| `AppContainerDown` | App container not running > 1m | critical |

Alerts are routed through Alertmanager to a Discord channel via `alertmanager-discord`.

---

## E2E Testing

End-to-end tests are written with [Playwright](https://playwright.dev/) and located in `go_app/tests/e2e/`.

```bash
cd go_app

# Install dependencies
npm install
npx playwright install

# Run tests (requires the app to be running on :8080)
npx playwright test
```

---

## Contributing

WhoKnows is not currently open for external contributions. The repository is public for presentation and reference purposes only.
