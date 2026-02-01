package memory

import (
	"context"

	"github.com/your-org/your-app/internal/ports"
)

// Consumer is an in-memory queue consumer for dev/testing.
type Consumer struct {
	enqueue chan *ports.Message
}

// NewConsumer creates an in-memory consumer.
func NewConsumer() *Consumer {
	return &Consumer{enqueue: make(chan *ports.Message, 64)}
}

// Subscribe blocks until ctx is done; processes messages from Enqueue.
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

// Enqueue pushes a message (for dev/testing).
func (c *Consumer) Enqueue(msg *ports.Message) {
	c.enqueue <- msg
}

var _ ports.Consumer = (*Consumer)(nil)
