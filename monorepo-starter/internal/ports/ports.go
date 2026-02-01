package ports

import (
	"context"

	"github.com/your-org/your-app/internal/domain"
)

// --- API ports (Hexagonal) ---

// Repository defines the persistence port. Implement in internal/adapters/persistence or pkg/database.
type Repository interface {
	GetByID(ctx context.Context, id string) (*domain.Entity, error)
	Save(ctx context.Context, e *domain.Entity) error
	Delete(ctx context.Context, id string) error
}

// Service defines the use case port for the API.
type Service interface {
	GetEntity(ctx context.Context, id string) (*domain.Entity, error)
	CreateEntity(ctx context.Context, e *domain.Entity) error
}

// --- Worker ports (queue) ---

// MessageHandler processes a single queue message.
type MessageHandler func(ctx context.Context, msg *Message) error

// Consumer is the queue consumer port. Implement in internal/adapters/queue or pkg/queue.
type Consumer interface {
	Subscribe(ctx context.Context, handler MessageHandler) error
}

// Message is the transport-agnostic queue message.
type Message struct {
	ID   string
	Body []byte
	Ack  func() error
	Nack func() error
}
