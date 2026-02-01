package service

import (
	"context"
	"time"

	"github.com/your-org/your-app/internal/domain"
	"github.com/your-org/your-app/internal/ports"
)

type svc struct {
	repo ports.Repository
}

// New creates the application service that uses the Repository port.
func New(repo ports.Repository) ports.Service {
	return &svc{repo: repo}
}

func (s *svc) GetEntity(ctx context.Context, id string) (*domain.Entity, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *svc) CreateEntity(ctx context.Context, e *domain.Entity) error {
	now := time.Now()
	e.CreatedAt = now
	e.UpdatedAt = now
	return s.repo.Save(ctx, e)
}
