package domain

import "errors"

var (
	// ErrNotFound is returned when a resource does not exist.
	ErrNotFound = errors.New("resource not found")
	// ErrInvalidInput is returned when the input is invalid.
	ErrInvalidInput = errors.New("invalid input")
	// ErrConflict is returned when there is a conflict (e.g. duplicate).
	ErrConflict = errors.New("conflict")
)
