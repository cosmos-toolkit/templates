package memory

import (
	"context"

	"github.com/your-org/your-worker/internal/ports"
)

// Consumer is an in-memory queue adapter for local dev and testing (plug-and-play).
// Replace with pkg/queue/rabbitmq, pkg/queue/sqs, etc. in production.
type Consumer struct {
	enqueue chan *ports.Message
}

// NewConsumer creates an in-memory consumer. For testing, push messages via Enqueue.
func NewConsumer() *Consumer {
	return &Consumer{enqueue: make(chan *ports.Message, 64)}
}

// Subscribe blocks until ctx is done. It processes messages from the in-memory queue.
// Use Enqueue from another goroutine to simulate incoming messages.
func (c *Consumer) Subscribe(ctx context.Context, handler ports.MessageHandler) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-c.enqueue:
			if err := handler(ctx, msg); err != nil && msg.Nack != nil {
				_ = msg.Nack()
			}
		}
	}
}

// Enqueue pushes a message to the consumer (for dev/testing).
func (c *Consumer) Enqueue(msg *ports.Message) {
	c.enqueue <- msg
}

// Ensure Consumer implements ports.Consumer at compile time.
var _ ports.Consumer = (*Consumer)(nil)
