package memory

import (
	"context"
	"sync"

	"github.com/your-org/your-app/internal/domain"
	"github.com/your-org/your-app/internal/ports"
)

// Repository is a memory adapter that implements ports.Repository.
type Repository struct {
	mu   sync.RWMutex
	data map[string]*domain.Entity
}

// NewRepository creates a memory repository.
func NewRepository() *Repository {
	return &Repository{data: make(map[string]*domain.Entity)}
}

func (r *Repository) GetByID(ctx context.Context, id string) (*domain.Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	e, ok := r.data[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return e, nil
}

func (r *Repository) Save(ctx context.Context, e *domain.Entity) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[e.ID] = e
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, id)
	return nil
}

var _ ports.Repository = (*Repository)(nil)
