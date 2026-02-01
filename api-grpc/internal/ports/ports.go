package ports

import (
	"context"

	"github.com/your-org/your-app/internal/domain"
)

// Repository defines the persistence port (driven adapter).
// Implementations are in internal/adapters/persistence or in pkg (plugins).
type Repository interface {
	GetByID(ctx context.Context, id string) (*domain.Entity, error)
	Save(ctx context.Context, e *domain.Entity) error
	Delete(ctx context.Context, id string) error
}

// Service defines the use case port (application calls the domain).
// Used by the gRPC adapter to orchestrate operations.
type Service interface {
	GetEntity(ctx context.Context, id string) (*domain.Entity, error)
	CreateEntity(ctx context.Context, e *domain.Entity) error
}
