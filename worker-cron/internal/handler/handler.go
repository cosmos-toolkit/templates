package handler

import (
	"context"
	"log"
	"time"
)

// Handler implements the job logic for each scheduled run. Replace with your domain use cases.
type Handler struct{}

// New creates a new job handler.
func New() *Handler {
	return &Handler{}
}

// Handle runs the scheduled job. Customize to call your application services (e.g. cleanup, sync, report).
func (h *Handler) Handle(ctx context.Context) error {
	log.Printf("handler: job run at %s", time.Now().UTC().Format(time.RFC3339))

	// Example: call use case, then return.
	// if err := h.useCase.Cleanup(ctx); err != nil { return err }
	// return nil

	return nil
}
