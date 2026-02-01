package handler

import (
	"context"
	"testing"

	"github.com/your-org/your-app/internal/ports"
)

func TestHandle_AcksOnSuccess(t *testing.T) {
	h := New()
	acked := false
	msg := &ports.Message{
		ID:   "1",
		Body: []byte("test"),
		Ack:  func() error { acked = true; return nil },
		Nack: func() error { return nil },
	}
	err := h.Handle(context.Background(), msg)
	if err != nil {
		t.Fatalf("Handle: %v", err)
	}
	if !acked {
		t.Error("expected Ack to be called")
	}
}
