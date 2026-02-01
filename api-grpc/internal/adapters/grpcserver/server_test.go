package grpcserver

import (
	"context"
	"testing"

	entityv1 "github.com/your-org/your-app/api/pb/entity/v1"
	"github.com/your-org/your-app/internal/app"
)

func TestGetEntity_NotFound(t *testing.T) {
	srv := NewServer(app.New())
	_, err := srv.GetEntity(context.Background(), &entityv1.GetEntityRequest{Id: "missing"})
	if err == nil {
		t.Error("expected error for missing entity")
	}
}

func TestCreateEntity(t *testing.T) {
	srv := NewServer(app.New())
	resp, err := srv.CreateEntity(context.Background(), &entityv1.CreateEntityRequest{Id: "e1"})
	if err != nil {
		t.Fatalf("CreateEntity: %v", err)
	}
	if resp.Id != "e1" {
		t.Errorf("got id %q", resp.Id)
	}
}
