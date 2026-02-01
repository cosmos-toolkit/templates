# Plugins (pkg)

Plugins are packages in `pkg/` that extend the API without changing the core (Hexagonal Architecture). Each plugin implements a **port** defined in `internal/ports` or provides utilities (config, logger, external clients).

## Conventions

- **Package name**: lowercase, single word when possible (`pkg/database`, `pkg/cache`, `pkg/env`).
- **Files**: `snake_case.go` for multiple words (e.g. `user_repository.go`).
- **Interface**: the plugin must implement an interface in `internal/ports` or expose a stable API for injection in `internal/app`.

## Suggested plugins

| Plugin         | Description                 | Usage in `app`                                 |
| -------------- | --------------------------- | ---------------------------------------------- |
| `pkg/env`      | Load variables (.env)       | Config in `cmd/api/main.go`                    |
| `pkg/database` | SQL client (Postgres, etc.) | `WithRepository(repo)` in `app.NewWithOptions` |
| `pkg/cache`    | Redis / in-memory cache     | New port `Cache` + `WithCache`                 |
| `pkg/auth`     | JWT, API Key, OAuth         | HTTP middleware in `router`                    |
| `pkg/queue`    | RabbitMQ, SQS, Kafka        | Consumers as adapters                          |
| `pkg/logger`   | Structured logger           | Injected in handlers and services              |

## Example: using a repository plugin

1. Create `pkg/database/postgres/repository.go` implementing `ports.Repository`.
2. In `cmd/api/main.go`:

```go
repo := postgres.NewRepository(cfg)
application := app.NewWithOptions(app.WithRepository(repo))
application.RegisterCloser(repo.Close)
```

The core (`internal/domain`, `internal/ports`, `internal/app/service`) remains free of external dependencies; only `cmd` and adapters import plugins.
