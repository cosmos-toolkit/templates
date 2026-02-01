package app

import (
	"context"
	"sync"

	"github.com/your-org/your-app/internal/adapters/persistence/memory"
	"github.com/your-org/your-app/internal/app/service"
	"github.com/your-org/your-app/internal/ports"
)

// App aggregates API dependencies (Hexagonal). Shared by cmd/api.
type App struct {
	Service ports.Service
	repo    ports.Repository
	closers []func(context.Context) error
	mu      sync.Mutex
}

// New creates the API application with default adapters.
func New() *App {
	repo := memory.NewRepository()
	svc := service.New(repo)
	return &App{
		Service: svc,
		repo:    repo,
		closers: nil,
	}
}

// NewWithOptions allows injecting port implementations (plugins).
func NewWithOptions(opts ...Option) *App {
	a := New()
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// Option configures the application.
type Option func(*App)

// WithRepository replaces the default repository (e.g. pkg/database).
func WithRepository(r ports.Repository) Option {
	return func(a *App) {
		a.repo = r
		a.Service = service.New(r)
	}
}

// RegisterCloser registers a shutdown callback (e.g. close DB).
func (a *App) RegisterCloser(fn func(context.Context) error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.closers = append(a.closers, fn)
}

// Shutdown runs all registered closers.
func (a *App) Shutdown(ctx context.Context) {
	a.mu.Lock()
	closers := append([]func(context.Context) error(nil), a.closers...)
	a.mu.Unlock()
	for _, fn := range closers {
		_ = fn(ctx)
	}
}
