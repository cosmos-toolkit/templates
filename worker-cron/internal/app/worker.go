package app

import (
	"context"
	"log"

	"github.com/your-org/your-worker/internal/ports"
)

// Worker runs a scheduled job with a scheduler (plug-and-play).
// Inject any Scheduler implementation (ticker, cron, cloud) via NewWorker.
type Worker struct {
	scheduler ports.Scheduler
	handler   ports.JobHandler
}

// NewWorker builds a worker with the given scheduler and job handler.
// Plug your scheduler: e.g. scheduler := cron.NewScheduler(cfg) from pkg/scheduler/cron.
func NewWorker(scheduler ports.Scheduler, handler ports.JobHandler) *Worker {
	return &Worker{scheduler: scheduler, handler: handler}
}

// Run starts the scheduler. It blocks until ctx is cancelled.
func (w *Worker) Run(ctx context.Context) error {
	log.Println("worker: starting scheduler")
	defer log.Println("worker: stopped")
	return w.scheduler.Run(ctx, w.handler)
}
