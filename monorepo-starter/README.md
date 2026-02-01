# Monorepo starter (Cosmos Template)

Single Go module with **API** (Hexagonal + chi), **Worker** (queue consumer), and **CLI** (Cobra). Shared `internal/` (domain, ports, app, adapters) and `pkg/`. Run everything locally or with **Docker Compose** (API + Worker).

## Project structure

```
monorepo-starter/
├── cmd/
│   ├── api/                 # HTTP API entrypoint
│   ├── worker/              # Queue worker entrypoint
│   └── cli/                 # CLI entrypoint
├── internal/
│   ├── domain/              # Shared entities and errors
│   ├── ports/               # API (Repository, Service) + Worker (Consumer, Message)
│   ├── app/                 # API composition (Hexagonal)
│   ├── adapters/
│   │   ├── http/            # API router and handlers (chi)
│   │   ├── persistence/     # API repository (memory)
│   │   └── queue/           # Worker consumer (memory)
│   ├── worker/              # Worker app and message handler
│   ├── cli/                 # CLI root command
│   └── commands/            # CLI subcommands (version, run)
├── pkg/
│   └── env/                 # Shared env helpers
├── configs/
├── build/
│   ├── Dockerfile.api
│   └── Dockerfile.worker
├── compose.yaml             # API + Worker (optional RabbitMQ)
├── go.mod
├── Makefile
└── README.md
```

## Quick start

### Local (no Docker)

```bash
# Terminal 1: API
make run-api
# or: go run ./cmd/api

# Terminal 2: Worker
make run-worker
# or: go run ./cmd/worker

# CLI
make build-cli && ./bin/cli version
make hello   # cli run --name "Cosmos"
```

### Docker Compose

```bash
make up
# API: http://localhost:8080/health
# Optional: add RabbitMQ in compose.yaml and set RABBITMQ_URL in worker
make down
```

## Endpoints (API)

| Method | Path                    |
|--------|-------------------------|
| GET    | /health                 |
| GET    | /api/v1/entities/{id}   |
| POST   | /api/v1/entities        |

## CLI commands

| Command   | Description       |
|----------|-------------------|
| version  | Print version     |
| run      | Sample (e.g. run -n "Cosmos") |

## Plug-and-play

- **API**: Replace `internal/adapters/persistence/memory` with `pkg/database` (e.g. Postgres) via `app.NewWithOptions(app.WithRepository(repo))` in `cmd/api/main.go`.
- **Worker**: Replace `internal/adapters/queue/memory` with `pkg/queue/rabbitmq` (or SQS) in `cmd/worker/main.go`; same `ports.Consumer` contract.
- **CLI**: Add commands in `internal/commands/` and register in `internal/cli/root.go`.

## Makefile targets

```bash
make build       # build api, worker, cli
make run-api     # run API
make run-worker  # run worker
make run-cli     # run CLI
make up          # docker compose up --build
make down        # docker compose down
make test        # run tests
make lint        # golangci-lint
```

## License

As per the Cosmos Toolkit project.
