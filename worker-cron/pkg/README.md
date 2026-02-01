# Scheduler plugins (plug-and-play)

Schedulers are **plugged in** by implementing the `ports.Scheduler` interface and swapping one line in `cmd/worker/main.go`. The core (ports, app, handler) does not depend on any specific cron or cloud implementation.

## Contract: `ports.Scheduler`

```go
type JobHandler func(ctx context.Context) error

type Scheduler interface {
    Run(ctx context.Context, handler JobHandler) error
}
```

Your adapter runs the handler on each tick/schedule (e.g. every minute, or at cron spec). When `ctx` is cancelled, `Run` should return.

## Suggested plugins

| Plugin                 | Description           | Usage in cmd/worker/main.go              |
|------------------------|-----------------------|------------------------------------------|
| `pkg/scheduler/cron`   | Cron expression (e.g. robfig/cron) | `scheduler := cron.NewScheduler("0 * * * *")` |
| `pkg/scheduler/cloud`  | GCP Cloud Scheduler + HTTP callback | Worker exposes HTTP; Cloud Scheduler hits it |
| `internal/adapters/scheduler/ticker` | In-memory ticker (dev) | Default in template; interval configurable |

## How to plug a scheduler

1. Add a package (e.g. `pkg/scheduler/cron`) that implements `ports.Scheduler`:
   - In `Run(ctx, handler)`, start a cron runner (or ticker) and call `handler(ctx)` on each tick.
   - When `ctx.Done()` is closed, stop the runner and return.

2. In `cmd/worker/main.go`, replace the default scheduler:

```go
// Default (ticker):
var scheduler ports.Scheduler = ticker.NewScheduler(1 * time.Minute)

// Plug cron:
scheduler := cron.NewScheduler(cron.Config{Spec: "0 * * * *"})  // every hour
```

3. Run the worker; no other code changes are required (plug-and-play).

## Example: minimal cron adapter (robfig/cron)

```go
// pkg/scheduler/cron/scheduler.go
package cron

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/your-org/your-worker/internal/ports"
)

type Scheduler struct {
	cron   *cron.Cron
	spec   string
}

func NewScheduler(spec string) *Scheduler {
	return &Scheduler{cron: cron.New(), spec: spec}
}

func (s *Scheduler) Run(ctx context.Context, handler ports.JobHandler) error {
	_, _ = s.cron.AddFunc(s.spec, func() {
		_ = handler(ctx)
	})
	s.cron.Start()
	<-ctx.Done()
	s.cron.Stop()
	return ctx.Err()
}
```

The worker binary stays the same; only the injected `Scheduler` implementation changes.
