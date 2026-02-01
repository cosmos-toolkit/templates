# API gRPC (Cosmos Template)

Go gRPC API template using **Hexagonal Architecture** (Ports & Adapters), with **buf** for proto and code generation.

## Project structure

```
api-grpc/
├── api/
│   ├── proto/                    # Proto definitions
│   │   └── entity/v1/entity.proto
│   └── pb/                        # Generated Go (run: make generate)
│       └── entity/v1/
├── cmd/
│   └── server/                    # gRPC server entrypoint
│       └── main.go
├── internal/
│   ├── domain/                    # Entities and errors
│   ├── ports/                     # Repository, Service interfaces
│   ├── app/                       # Composition and use cases
│   └── adapters/
│       ├── grpcserver/            # Implements EntityServiceServer
│       └── persistence/memory/    # In-memory repository
├── pkg/                           # Optional plugins (env, database, etc.)
├── configs/
├── docs/
│   └── ARCHITECTURE.md
├── buf.yaml
├── buf.gen.yaml
├── go.mod
├── Makefile
└── README.md
```

## Prerequisites

- Go 1.22+
- [buf](https://buf.build/docs/installation) (to regenerate code from proto)

The template ships with pre-generated `api/pb/` so it builds without buf. After editing `.proto` files, run:

```bash
make generate
```

## How to use

1. Copy the template and set the `module` in `go.mod` (e.g. `github.com/your-org/your-grpc-api`).
2. Replace `internal/domain` with your entities and errors.
3. Extend `api/proto` and run `make generate`; implement or extend the gRPC adapter in `internal/adapters/grpcserver`.
4. Run:

   ```bash
   make run
   # or: go run ./cmd/server
   ```

   Server listens on `:9090` (override with `GRPC_ADDR`).

## gRPC methods

| Method       | Request     | Response |
| ------------ | ----------- | -------- |
| GetEntity    | id (string) | Entity   |
| CreateEntity | id (string) | Entity   |

Reflection is enabled for tools like `grpcurl` and Evans.

## Build and tests

```bash
make build   # bin/server
make run     # start the gRPC server
make test    # run tests
make lint    # golangci-lint (requires installation)
make generate # regenerate from proto (requires buf)
```

## Plugins (pkg)

Same pattern as the HTTP API template: implement `ports.Repository` in `pkg/database` and inject with `app.NewWithOptions(app.WithRepository(repo))`. See `pkg/README.md`.

## License

As per the Cosmos Toolkit project.
