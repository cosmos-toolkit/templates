# Queue worker (Cosmos Template)

Go worker template that **consumes messages from a queue** with a **plug-and-play** design: the queue client is a pluggable adapter implementing `ports.Consumer`; swap it in `cmd/worker/main.go` (e.g. memory → RabbitMQ, SQS, Kafka) without changing app or handler code.

## Project structure

```
worker-queue/
├── cmd/
│   └── worker/              # Entrypoint: wire Consumer, run Worker
│       └── main.go
├── internal/
│   ├── ports/                # Consumer interface, Message, MessageHandler
│   │   └── consumer.go
│   ├── app/                  # Worker (runs consumer with handler)
│   │   └── worker.go
│   ├── handler/              # Message processing logic (your use cases)
│   │   └── handler.go
│   ├── domain/               # Domain errors (optional)
│   │   └── errors.go
│   └── adapters/
│       └── queue/            # Queue adapters (implement Consumer)
│           └── memory/       # In-memory consumer for dev/testing
├── pkg/                      # Queue plugins (rabbitmq, sqs, kafka, etc.)
│   └── README.md             # How to implement and plug a Consumer
├── configs/
├── docs/
│   └── ARCHITECTURE.md
├── go.mod
├── Makefile
└── README.md
```

## Plug-and-play

- **Port**: `ports.Consumer` — `Subscribe(ctx, handler) error`; handler receives `*ports.Message` (ID, Body, Ack, Nack).
- **Default adapter**: `internal/adapters/queue/memory` — in-memory consumer; use for local dev or tests.
- **Production**: implement `ports.Consumer` in `pkg/queue/rabbitmq` (or sqs, kafka) and in `cmd/worker/main.go` set:
  ```go
  consumer := rabbitmq.NewConsumer(rabbitmq.Config{...})
  ```
  No other code changes needed.

See `pkg/README.md` for the contract and an example RabbitMQ adapter.

## How to use

1. Copy the template and set the `module` in `go.mod` (e.g. `github.com/your-org/your-worker`).
2. Replace or extend `internal/handler/handler.go` with your message processing logic (parse body, call services, ack/nack).
3. For production, add a queue plugin (e.g. `pkg/queue/rabbitmq`) implementing `ports.Consumer` and use it in `cmd/worker/main.go`.
4. Run:
   ```bash
   make run
   # or: go run ./cmd/worker
   ```
   With the default memory consumer, the worker runs until SIGINT/SIGTERM; use `Enqueue` from tests or another goroutine to push messages.

## Build and tests

```bash
make build   # bin/worker
make run     # run worker
make test    # run tests
make lint    # golangci-lint (requires installation)
```

## License

As per the Cosmos Toolkit project.
