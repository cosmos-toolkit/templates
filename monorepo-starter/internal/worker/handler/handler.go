package handler

import (
	"context"
	"log"

	"github.com/your-org/your-app/internal/ports"
)

// Handler processes queue messages.
type Handler struct{}

// New creates the message handler.
func New() *Handler {
	return &Handler{}
}

// Handle processes a single message. Replace with your use cases.
func (h *Handler) Handle(ctx context.Context, msg *ports.Message) error {
	log.Printf("worker: message id=%s body=%s", msg.ID, string(msg.Body))
	return msg.Ack()
}
