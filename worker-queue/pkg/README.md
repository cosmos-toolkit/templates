# Queue plugins (plug-and-play)

Queue clients are **plugged in** by implementing the `ports.Consumer` interface and swapping one line in `cmd/worker/main.go`. The core (ports, app, handler) does not depend on any specific broker.

## Contract: `ports.Consumer`

```go
type MessageHandler func(ctx context.Context, msg *ports.Message) error

type Consumer interface {
    Subscribe(ctx context.Context, handler MessageHandler) error
}
```

`ports.Message` has `ID`, `Body []byte`, `Ack func() error`, `Nack func() error`. Your adapter wraps the broker’s native message and calls the handler for each delivery; the handler calls `Ack()` or `Nack()` when done.

## Suggested plugins

| Plugin                           | Description       | Usage in cmd/worker/main.go             |
| -------------------------------- | ----------------- | --------------------------------------- |
| `pkg/queue/rabbitmq`             | RabbitMQ consumer | `consumer := rabbitmq.NewConsumer(cfg)` |
| `pkg/queue/sqs`                  | AWS SQS consumer  | `consumer := sqs.NewConsumer(cfg)`      |
| `pkg/queue/kafka`                | Kafka consumer    | `consumer := kafka.NewConsumer(cfg)`    |
| `internal/adapters/queue/memory` | In-memory (dev)   | Default in template; no extra pkg       |

## How to plug a queue client

1. Add a package (e.g. `pkg/queue/rabbitmq`) that implements `ports.Consumer`:

   - Connect to the broker in `NewConsumer(cfg)`.
   - In `Subscribe(ctx, handler)`, start consuming and for each message:
     - Build a `ports.Message` with `ID`, `Body`, and `Ack`/`Nack` that call the broker’s ack/nack.
     - Call `handler(ctx, msg)`.
   - When `ctx` is done, stop consuming and return.

2. In `cmd/worker/main.go`, replace the default consumer:

```go
// Default (in-memory):
var consumer ports.Consumer = memory.NewConsumer()

// Plug RabbitMQ:
consumer := rabbitmq.NewConsumer(rabbitmq.Config{
    URL:   os.Getenv("RABBITMQ_URL"),
    Queue: "tasks",
})
```

3. Run the worker; no other code changes are required (plug-and-play).

## Example: minimal RabbitMQ adapter

```go
// pkg/queue/rabbitmq/consumer.go
package rabbitmq

import (
    "context"
    "github.com/your-org/your-worker/internal/ports"
    amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct{ ch *amqp.Channel; msgs <-chan amqp.Delivery }

func NewConsumer(cfg Config) (*Consumer, error) {
    // ... connect, declare queue, ch.Consume(...)
    return &Consumer{ch: ch, msgs: deliveries}, nil
}

func (c *Consumer) Subscribe(ctx context.Context, handler ports.MessageHandler) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case d := <-c.msgs:
            msg := &ports.Message{
                ID: d.MessageId, Body: d.Body,
                Ack:  func() error { return d.Ack(false) },
                Nack: func() error { return d.Nack(false, true) },
            }
            if err := handler(ctx, msg); err != nil && msg.Nack != nil {
                _ = msg.Nack()
            }
        }
    }
}
```

The worker binary stays the same; only the injected `Consumer` implementation changes.
