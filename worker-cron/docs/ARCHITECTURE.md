# Cron / scheduled worker architecture

This template runs a **worker** that executes a job on a **schedule** (ticker, cron, or cloud). The scheduler is **pluggable**: the core depends only on the `ports.Scheduler` interface; you swap the implementation (ticker, cron, cloud) in `cmd/worker/main.go` without changing app or handler code.

## Layers

```
                    +------------------+
                    | cmd/worker       |  Wiring: choose Scheduler, run Worker
                    +--------+---------+
                             |
                    +--------v---------+
                    | internal/ports   |  Scheduler interface, JobHandler
                    +--------+---------+
                             |
    +------------------------+------------------------+
    |                        |                        |
+---v---+              +-----v-----+              +----v----+
| handler|              | app/worker|              | adapters |
|        |              |           |              | (scheduler)|
+---+---+              +-----+-----+              +----+----+
    |                        |                        |
    |                        |              ticker, or pkg/scheduler/*
    |                        |              (cron, cloud)
    +------------------------+------------------------+
```

### Ports (`internal/ports`)

- **Scheduler**: `Run(ctx, handler) error` — blocks until ctx is done; calls `handler` on each tick/schedule.
- **JobHandler**: `func(ctx) error` — the work to run on each schedule (implemented by `internal/handler`).

### App (`internal/app`)

- **Worker**: holds a `Scheduler` and a `JobHandler`; `Run(ctx)` calls `scheduler.Run(ctx, handler)`.
- No dependency on any specific scheduler; only on `ports.Scheduler`.

### Handler (`internal/handler`)

- Implements the job logic for each run (cleanup, sync, report, etc.).
- Replace or extend with your domain use cases.

### Adapters (scheduler)

- **internal/adapters/scheduler/ticker**: in-memory ticker for local dev and tests; runs handler every N duration.
- **pkg/scheduler/***: production implementations (e.g. `pkg/scheduler/cron`) that implement `ports.Scheduler` and are injected in `cmd/worker/main.go`.

## Plug-and-play flow

1. **Default**: `cmd/worker/main.go` uses `ticker.NewScheduler(1 * time.Minute)`; worker runs the job every minute until ctx is cancelled.
2. **Plug cron/cloud**: add `pkg/scheduler/cron`, implement `Scheduler`, then in main: `scheduler := cron.NewScheduler(cfg)` and pass it to `app.NewWorker(scheduler, handler.Handle)`.
3. Same binary layout; only the constructor and dependency in main change.

## Naming (Go)

- **Directories**: lowercase, single word (`ports`, `app`, `handler`, `adapters`).
- **Files**: `snake_case.go` for multiple words (`job_handler.go`).
- **Interface**: defined where it is consumed (`ports.Scheduler`); implemented in adapters or pkg.
