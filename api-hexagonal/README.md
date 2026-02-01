# API Hexagonal (Cosmos Template)

Go API template using **Hexagonal Architecture** (Ports & Adapters), following Go community conventions and ready for extension via **plugins** in `pkg/`.

## Project structure

```
api-hexagonal/
├── cmd/
│   └── api/                 # Application entrypoint
│       └── main.go
├── internal/                # Private application code
│   ├── domain/              # Domain entities and errors
│   ├── ports/               # Interfaces (Repository, Service, etc.)
│   ├── app/                 # Composition and dependency injection
│   │   └── service/         # Use cases (application)
│   └── adapters/            # Port implementations
│       ├── http/            # HTTP handlers and router
│       │   ├── handler/
│       │   └── router/
│       └── persistence/     # Repositories (e.g. memory, then plugins)
│           └── memory/
├── pkg/                     # Reusable plugins (database, cache, auth, env, queue, ...)
│   └── README.md            # How to create and use plugins
├── configs/                 # Example configuration
├── build/                   # Docker, CI (optional)
├── go.mod
├── Makefile
└── README.md
```

## Conventions (Go community)

- **Directories**: lowercase, single word when possible (`domain`, `ports`, `adapters`).
- **Files**: `snake_case.go` for compound names (e.g. `user_repository.go`).
- **Packages**: directory name = package name; short and clear.
- **Ports**: interfaces in `internal/ports`; adapters (including plugins in `pkg/`) implement them.
- **Entrypoint**: minimal `cmd/api/main.go`; all logic in `internal` and `pkg`.

## How to use

1. **Copy the template** to your repo and set the `module` in `go.mod` (e.g. `github.com/your-org/your-api`).
2. **Replace** `internal/domain/entity.go` and ports with your entities and contracts.
3. **Run**:
   ```bash
   make run
   # or: go run ./cmd/api
   ```
4. **Try the endpoints**:
   - `GET /health`
   - `GET /api/v1/entities/{id}`
   - `POST /api/v1/entities` (body: `{"id": "123"}`)

## Plugins (pkg)

Plugins are packages in `pkg/` that implement ports or provide config, logger, DB/cache/queue clients. The core (`internal/domain`, `internal/ports`, `internal/app/service`) does not depend on them; composition is done in `cmd/api/main.go` with `app.NewWithOptions(...)`.

Example plugins:

- **pkg/env** – environment variables / .env
- **pkg/database** – Postgres repository (implements `ports.Repository`)
- **pkg/cache** – Redis or similar (add a `Cache` port if needed)
- **pkg/auth** – JWT / API Key (middleware in router)
- **pkg/queue** – queue clients (consumers as adapters)
- **pkg/logger** – structured logger

See `pkg/README.md` for conventions and usage examples.

## Build and tests

```bash
make build   # bin/api
make run     # start the API
make test    # run tests
make lint    # golangci-lint (requires installation)
```

## License

As per the Cosmos Toolkit project.
