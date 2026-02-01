# Monorepo architecture

Single module with three entrypoints (API, Worker, CLI) sharing `internal/` and `pkg/`.

## Layout

```
cmd/api     → internal/app (Hexagonal) + internal/adapters/http, persistence
cmd/worker  → internal/worker + internal/adapters/queue + ports.Consumer
cmd/cli     → internal/cli + internal/commands
```

## Shared packages

- **internal/domain**: Entities and errors (used by API and optionally worker).
- **internal/ports**: API ports (Repository, Service) and Worker ports (Consumer, Message). Same package; different interfaces.
- **pkg/env**: Env helpers (used by any cmd).

## API (Hexagonal)

- **internal/app**: Composition; `app.New()` or `app.NewWithOptions(WithRepository(...))`.
- **internal/app/service**: Use case (GetEntity, CreateEntity).
- **internal/adapters/http**: Chi router and handlers.
- **internal/adapters/persistence**: Memory repository; replace with pkg/database in production.

## Worker (queue)

- **internal/worker**: `worker.New(consumer, handler)` and `Run(ctx)`.
- **internal/worker/handler**: Message handler (process message, ack/nack).
- **internal/adapters/queue**: Memory consumer; replace with pkg/queue in production.

## CLI

- **internal/cli**: Root Cobra command; registers subcommands.
- **internal/commands**: version, run (add more as needed).

## Docker Compose

- **api**: Builds from `build/Dockerfile.api`, exposes 8080.
- **worker**: Builds from `build/Dockerfile.worker`; optional env for RabbitMQ.
- **rabbitmq**: Optional service; uncomment in compose.yaml and set worker env to use it.

## Adding a new app

1. Add `cmd/myapp/main.go` and wire shared internal packages.
2. Optionally add `build/Dockerfile.myapp` and a service in compose.yaml.
