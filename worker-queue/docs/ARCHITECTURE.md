# Plug-and-play worker architecture

This template runs a **worker** that consumes messages from a queue. The queue client is **pluggable**: the core depends only on the `ports.Consumer` interface; you swap the implementation (memory, RabbitMQ, SQS, Kafka) in `cmd/worker/main.go` without changing app or handler code.

## Layers

```
                    +------------------+
                    | cmd/worker       |  Wiring: choose Consumer, run Worker
                    +--------+---------+
                             |
                    +--------v---------+
                    | internal/ports   |  Consumer interface, Message, MessageHandler
                    +--------+---------+
                             |
    +------------------------+------------------------+
    |                        |                        |
+---v---+              +-----v-----+              +----v----+
| handler|              | app/worker|              | adapters |
|        |              |           |              | (queue)  |
+---+---+              +-----+-----+              +----+----+
    |                        |                        |
    |                        |              memory, or pkg/queue/*
    |                        |              (rabbitmq, sqs, kafka)
    +------------------------+------------------------+
```

### Ports (`internal/ports`)

- **Consumer**: `Subscribe(ctx, handler) error` — blocks until ctx is done; calls `handler` for each message.
- **Message**: `ID`, `Body`, `Ack`, `Nack` — transport-agnostic; adapters wrap broker messages into this type.
- **MessageHandler**: `func(ctx, msg) error` — implemented by `internal/handler`; call `msg.Ack()` or `msg.Nack()` inside.

### App (`internal/app`)

- **Worker**: holds a `Consumer` and a `MessageHandler`; `Run(ctx)` calls `consumer.Subscribe(ctx, handler)`.
- No dependency on any specific queue; only on `ports.Consumer`.

### Handler (`internal/handler`)

- Implements the business logic for each message (parse body, call use cases, then ack or nack).
- Replace or extend with your domain logic.

### Adapters (queue)

- **internal/adapters/queue/memory**: in-memory consumer for local dev and tests; messages are pushed via `Enqueue`.
- **pkg/queue/\***: production implementations (e.g. `pkg/queue/rabbitmq`) that implement `ports.Consumer` and are injected in `cmd/worker/main.go`.

## Plug-and-play flow

1. **Default**: `cmd/worker/main.go` uses `memory.NewConsumer()`; worker runs and waits for messages (or ctx cancel).
2. **Plug a broker**: add `pkg/queue/rabbitmq`, implement `Consumer`, then in main: `consumer := rabbitmq.NewConsumer(cfg)` and pass it to `app.NewWorker(consumer, handler.Handle)`.
3. Same binary layout; only the constructor and dependency in main change.

## Naming (Go)

- **Directories**: lowercase, single word (`ports`, `app`, `handler`, `adapters`).
- **Files**: `snake_case.go` for multiple words (`message_handler.go`).
- **Interface**: defined where it is consumed (`ports.Consumer`); implemented in adapters or pkg.
