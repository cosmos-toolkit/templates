// Package sdk provides a client for the Your API.
// Use NewClient to create a client and call its methods.
package sdk

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Client is the main entrypoint for the SDK. It is safe for concurrent use.
type Client struct {
	baseURL    string
	httpClient *http.Client
	apiKey     string
}

// Option configures a Client.
type Option func(*Client)

// WithBaseURL sets the API base URL (default: https://api.example.com).
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithHTTPClient sets the HTTP client (default: http.DefaultClient with 30s timeout).
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient = client
	}
}

// WithAPIKey sets the API key for authenticated requests.
func WithAPIKey(key string) Option {
	return func(c *Client) {
		c.apiKey = key
	}
}

// NewClient creates a new SDK client. Pass options to customize (base URL, HTTP client, API key).
func NewClient(opts ...Option) *Client {
	c := &Client{
		baseURL: "https://api.example.com",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Ping checks connectivity to the API. Replace with your first real method.
func (c *Client) Ping(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/health", nil)
	if err != nil {
		return err
	}
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ping: status %d", resp.StatusCode)
	}
	return nil
}

// GetResource fetches a resource by ID. Replace with your domain types and logic.
func (c *Client) GetResource(ctx context.Context, id string) (*Resource, error) {
	if id == "" {
		return nil, fmt.Errorf("id required")
	}
	// Example: build request, do HTTP call, decode response.
	_ = ctx
	return &Resource{ID: id, Name: "example"}, nil
}
