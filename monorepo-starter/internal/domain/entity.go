package domain

import "time"

// Entity represents a domain entity (shared by API and worker if needed).
// Replace with your own entities (User, Order, etc.).
type Entity struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
