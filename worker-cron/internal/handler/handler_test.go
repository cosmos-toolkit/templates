package handler

import (
	"context"
	"testing"
)

func TestHandle_Succeeds(t *testing.T) {
	h := New()
	err := h.Handle(context.Background())
	if err != nil {
		t.Fatalf("Handle: %v", err)
	}
}
