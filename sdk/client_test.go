package sdk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient()
	if c == nil {
		t.Fatal("NewClient returned nil")
	}
	if c.baseURL != "https://api.example.com" {
		t.Errorf("default baseURL: got %q", c.baseURL)
	}
}

func TestNewClientWithOptions(t *testing.T) {
	url := "https://custom.example.com"
	c := NewClient(WithBaseURL(url), WithAPIKey("test-key"))
	if c.baseURL != url {
		t.Errorf("baseURL: got %q", c.baseURL)
	}
	if c.apiKey != "test-key" {
		t.Errorf("apiKey: got %q", c.apiKey)
	}
}

func TestClient_Ping(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/health" {
			t.Errorf("path: got %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := NewClient(WithBaseURL(server.URL))
	if err := c.Ping(context.Background()); err != nil {
		t.Errorf("Ping: %v", err)
	}
}

func TestClient_GetResource(t *testing.T) {
	c := NewClient()
	res, err := c.GetResource(context.Background(), "id1")
	if err != nil {
		t.Fatalf("GetResource: %v", err)
	}
	if res.ID != "id1" {
		t.Errorf("ID: got %q", res.ID)
	}
}

func TestClient_GetResource_EmptyID(t *testing.T) {
	c := NewClient()
	_, err := c.GetResource(context.Background(), "")
	if err == nil {
		t.Error("expected error for empty id")
	}
}
