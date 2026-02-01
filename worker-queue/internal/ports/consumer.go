package ports

import (
	"context"
)

// MessageHandler processes a single message. Return nil to ack, non-nil to nack (or use msg.Nack() explicitly).
// The handler may call msg.Ack() or msg.Nack() itself for finer control (e.g. ack after DB commit).
type MessageHandler func(ctx context.Context, msg *Message) error

// Consumer is the port for queue consumption (plug-and-play).
// Implement this interface in pkg/queue (e.g. pkg/queue/rabbitmq, pkg/queue/sqs) and inject it in cmd/worker.
// Subscribe blocks until ctx is cancelled; it receives messages and calls handler for each.
type Consumer interface {
	Subscribe(ctx context.Context, handler MessageHandler) error
}

// Message is the transport-agnostic message type used by Consumer and MessageHandler.
// It is defined here so that both the port and adapters (pkg/queue/*) use the same contract.
type Message struct {
	ID   string
	Body []byte
	Ack  func() error
	Nack func() error
}
