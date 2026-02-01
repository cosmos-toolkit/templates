package handler

import (
	"context"
	"log"

	"github.com/your-org/your-worker/internal/ports"
)

// Handler implements message processing logic. Replace with your domain use cases.
type Handler struct{}

// New creates a new message handler.
func New() *Handler {
	return &Handler{}
}

// Handle processes a single message. Ack on success, Nack on failure.
// Customize this method to call your application services (e.g. send email, update DB).
func (h *Handler) Handle(ctx context.Context, msg *ports.Message) error {
	log.Printf("handler: received message id=%s body=%s", msg.ID, string(msg.Body))

	// Example: parse body, call use case, then ack or nack.
	// if err := h.useCase.Process(ctx, msg.Body); err != nil { return msg.Nack() }
	// return msg.Ack()

	return msg.Ack()
}
