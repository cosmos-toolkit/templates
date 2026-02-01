# gRPC API (Hexagonal)

This template organizes the gRPC server using **Ports & Adapters** (Hexagonal), aligned with Go community conventions.

## Layers

```
                    +------------------+
                    |  gRPC adapter    |  (Driving adapter)
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
    |              | persistence, grpc |<-------------+
    +--------------+------------------+
```

### Domain (`internal/domain`)

- Domain entities and errors. No dependencies on frameworks or proto.

### Ports (`internal/ports`)

- **Repository** (driven): persistence.
- **Service** (driving): use cases used by the gRPC adapter.

### Application (`internal/app`)

- Composition and use cases. Same pattern as the HTTP API template.

### Adapters (`internal/adapters`)

- **grpcserver**: implements the generated `EntityServiceServer` interface, delegates to `app.Service`, converts proto ↔ domain.
- **persistence/memory**: in-memory repository for development.

### Proto and generated code (`api/`)

- `api/proto/`: `.proto` definitions.
- `api/pb/`: generated Go code (run `make generate` with buf to regenerate).

## Request flow

1. Client calls gRPC (e.g. `GetEntity`).
2. `internal/adapters/grpcserver` receives the request, converts proto to domain, calls `app.Service`.
3. Service uses `ports.Repository` and domain types.
4. Adapter converts domain result to proto and returns.

## Regenerating from proto

After editing `.proto` files, run:

```bash
make generate
```

Requires [buf](https://buf.build/docs/installation).
