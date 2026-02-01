package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/your-org/your-app/internal/adapters/http/handler"
	"github.com/your-org/your-app/internal/app"
)

// New returns an http.Handler with the API routes (chi router).
func New(a *app.App) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := handler.New(a)

	r.Get("/health", h.Health)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/entities/{id}", h.GetEntity)
		r.Post("/entities", h.CreateEntity)
	})

	return r
}
