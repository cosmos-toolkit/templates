package domain

import "errors"

var (
	// ErrInvalidPayload is returned when the message body cannot be processed.
	ErrInvalidPayload = errors.New("invalid payload")
	// ErrRetryLater is returned when the handler wants to nack and retry.
	ErrRetryLater = errors.New("retry later")
)
