package domain

import "time"

// Entity represents an aggregate or domain entity.
// Replace with your entities (User, Order, etc.).
type Entity struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
