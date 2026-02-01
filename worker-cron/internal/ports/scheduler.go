package ports

import "context"

// JobHandler runs a single scheduled job. Return nil on success, non-nil on failure.
// The scheduler may retry or log depending on the adapter (cron, cloud, ticker).
type JobHandler func(ctx context.Context) error

// Scheduler is the port for time-based execution (plug-and-play).
// Implement this interface in pkg/scheduler (e.g. pkg/scheduler/cron, pkg/scheduler/cloud) and inject it in cmd/worker.
// Run blocks until ctx is cancelled; it invokes handler on each tick/schedule.
type Scheduler interface {
	Run(ctx context.Context, handler JobHandler) error
}
