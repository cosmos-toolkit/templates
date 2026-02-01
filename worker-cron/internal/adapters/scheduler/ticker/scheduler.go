package ticker

import (
	"context"
	"time"

	"github.com/your-org/your-worker/internal/ports"
)

// Scheduler is an in-memory ticker adapter for local dev and testing (plug-and-play).
// Replace with pkg/scheduler/cron, pkg/scheduler/cloud, etc. in production.
type Scheduler struct {
	interval time.Duration
}

// NewScheduler creates a ticker-based scheduler that runs the job every interval.
// Example: NewScheduler(30 * time.Second) for every 30 seconds.
func NewScheduler(interval time.Duration) *Scheduler {
	if interval <= 0 {
		interval = 1 * time.Minute
	}
	return &Scheduler{interval: interval}
}

// Run blocks until ctx is done. It runs the handler on each tick.
func (s *Scheduler) Run(ctx context.Context, handler ports.JobHandler) error {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	// Run once immediately, then on each tick
	if err := handler(ctx); err != nil {
		// Log but continue; adapter may be configured to retry
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := handler(ctx); err != nil {
				// Log but continue
			}
		}
	}
}

// Ensure Scheduler implements ports.Scheduler at compile time.
var _ ports.Scheduler = (*Scheduler)(nil)
