package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/your-org/your-app/internal/app"
)

func TestHealth(t *testing.T) {
	h := New(app.New())
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	h.Health(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Health: got status %d", rec.Code)
	}
}
