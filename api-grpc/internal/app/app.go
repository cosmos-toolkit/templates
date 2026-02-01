package app

import (
	"context"
	"sync"

	"github.com/your-org/your-app/internal/adapters/persistence/memory"
	"github.com/your-org/your-app/internal/app/service"
	"github.com/your-org/your-app/internal/ports"
)

// App aggregates dependencies and exposes ports for the gRPC adapter.
// Plugins (pkg/database, pkg/cache, etc.) can be injected here.
type App struct {
	Service ports.Service
	repo    ports.Repository
	closers []func(context.Context) error
	mu      sync.Mutex
}

// New creates the application with default adapters.
func New() *App {
	repo := memory.NewRepository()
	svc := service.New(repo)
	return &App{
		Service: svc,
		repo:    repo,
		closers: nil,
	}
}

// NewWithOptions allows injecting implementations of ports (plugins).
func NewWithOptions(opts ...Option) *App {
	a := New()
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// Option configures the application.
type Option func(*App)

// WithRepository replaces the default repository (e.g. plugin pkg/database).
func WithRepository(r ports.Repository) Option {
	return func(a *App) {
		a.repo = r
		a.Service = service.New(r)
	}
}

// RegisterCloser registers a callback for shutdown.
func (a *App) RegisterCloser(fn func(context.Context) error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.closers = append(a.closers, fn)
}

// Shutdown executes all registered closers.
func (a *App) Shutdown(ctx context.Context) {
	a.mu.Lock()
	closers := append([]func(context.Context) error(nil), a.closers...)
	a.mu.Unlock()
	for _, fn := range closers {
		_ = fn(ctx)
	}
}
