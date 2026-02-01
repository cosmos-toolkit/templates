# Hexagonal Architecture

This template organizes code using **Ports & Adapters** (Hexagonal), aligned with Go community conventions.

## Layers

```
                    +------------------+
                    |   HTTP / CLI     |  (Driving adapters)
                    +--------+---------+
                             |
                    +--------v---------+
                    |   internal/ports |  (Interfaces)
                    +--------+---------+
                             |
    +------------------------+------------------------+
    |                        |                        |
+---v---+              +-----v-----+              +----v----+
| domain|              | app/service|              | plugins |
|       |              | (use cases) |              | (pkg/)  |
+---+---+              +-----+-----+              +----+----+
    |                        |                        |
    |              +---------v---------+              |
    |              | internal/adapters |  (Driven)    |
    |              | persistence, http |<-------------+
    +--------------+------------------+
```

### Domain (`internal/domain`)

- Domain entities, value objects, and errors.
- **No dependencies** on frameworks or infrastructure.
- Files: `entity.go`, `errors.go`, `value_object.go` (if needed).

### Ports (`internal/ports`)

- **Interfaces** the core exposes (driving) or consumes (driven).
- E.g. `Repository` (driven), `Service` (driving – use case).
- Plugins in `pkg/` implement ports (e.g. `pkg/database` implements `Repository`).

### Application (`internal/app`)

- **Composition**: `app.New()` or `app.NewWithOptions(WithRepository(...))`.
- **Use cases**: `internal/app/service` orchestrates domain and ports.
- Registers closers for shutdown (plugins that need to close connections).

### Adapters (`internal/adapters`)

- **Driving**: HTTP (handlers, router), later CLI or gRPC.
- **Driven**: persistence (e.g. `memory`), external clients; may come from `pkg/` (plugins).

### Plugins (`pkg/`)

- Optional packages that implement ports or provide config/logger/queue.
- Not imported by domain or ports; only by `cmd` and adapters.
- E.g. `pkg/database`, `pkg/cache`, `pkg/env`, `pkg/auth`, `pkg/queue`, `pkg/logger`.

## Request flow

1. `cmd/api/main.go` → creates `app` (with or without plugins).
2. `router` sets up routes and calls **handlers**.
3. Handler calls `app.Service` (port).
4. Service uses **Repository** (port) and **domain**.
5. Repository may be in-memory (`internal/adapters`) or a plugin (`pkg/database`).

## Naming (Go)

- **Directories**: single name, lowercase (`domain`, `ports`, `adapters`).
- **Files**: `snake_case.go` for multiple words (`user_repository.go`).
- **Interfaces**: noun or verb + “er” for behaviour (`Repository`, `Writer`).
- **Implementations**: optional suffix by type (`memory.Repository`, `postgres.Repository`).
