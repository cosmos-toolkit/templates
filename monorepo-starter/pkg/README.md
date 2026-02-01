# Shared packages (pkg)

Packages in `pkg/` are shared by **cmd/api**, **cmd/worker**, and **cmd/cli**. The core (internal/domain, internal/ports, internal/app) does not depend on pkg; only cmd and adapters import them.

## Conventions

- **pkg/env** – environment variables / .env (used by any cmd or config).
- **pkg/database** – DB client; implement `ports.Repository` in internal/adapters/persistence/postgres using pkg/database, then inject in cmd/api with `app.NewWithOptions(app.WithRepository(repo))`.
- **pkg/queue** – queue client (RabbitMQ, SQS); implement `ports.Consumer` and inject in cmd/worker.
- **pkg/logger** – structured logger (inject in handlers or worker).

## Plug-and-play

- **API**: replace internal/adapters/persistence/memory with a repository that uses pkg/database; in cmd/api: `app.NewWithOptions(app.WithRepository(postgres.NewRepository(cfg)))`.
- **Worker**: replace internal/adapters/queue/memory with pkg/queue/rabbitmq; in cmd/worker: `consumer := rabbitmq.NewConsumer(cfg)`.
- **CLI**: add commands in internal/commands or pkg/cli/commands and register in internal/cli/root.go.

Commands and entrypoints stay thin; shared logic lives in internal or pkg.
