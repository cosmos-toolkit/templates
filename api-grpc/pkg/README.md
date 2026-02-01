# Plugins (pkg)

Plugins extend the gRPC API without changing the core. Implement ports in `internal/ports` or provide utilities (config, logger, database).

## Suggested plugins

| Plugin         | Description           | Usage in `app`                  |
| -------------- | --------------------- | ------------------------------- |
| `pkg/env`      | Environment variables | Config in `cmd/server/main.go`  |
| `pkg/database` | Postgres, etc.        | `WithRepository(repo)` in `app` |
| `pkg/logger`   | Structured logger     | Injected in adapters            |

## Example: Postgres repository

1. Create `pkg/database/postgres/repository.go` implementing `ports.Repository`.
2. In `cmd/server/main.go`: `app.NewWithOptions(app.WithRepository(repo))`.
