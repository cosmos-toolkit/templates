# Cron / scheduled worker (Cosmos Template)

Go worker template that **runs a job on a schedule** (ticker, cron, or cloud) with a **plug-and-play** design: the scheduler is a pluggable adapter implementing `ports.Scheduler`; swap it in `cmd/worker/main.go` (e.g. ticker → cron, Cloud Scheduler) without changing app or handler code.

## Project structure

```
worker-cron/
├── cmd/
│   └── worker/              # Entrypoint: wire Scheduler, run Worker
│       └── main.go
├── internal/
│   ├── ports/                # Scheduler interface, JobHandler
│   │   └── scheduler.go
│   ├── app/                  # Worker (runs scheduler with handler)
│   │   └── worker.go
│   ├── handler/              # Job logic (your use cases)
│   │   └── handler.go
│   └── adapters/
│       └── scheduler/        # Scheduler adapters (implement Scheduler)
│           └── ticker/       # In-memory ticker for dev/testing
├── pkg/                      # Scheduler plugins (cron, cloud, etc.)
│   └── README.md             # How to implement and plug a Scheduler
├── configs/
├── docs/
│   └── ARCHITECTURE.md
├── go.mod
├── Makefile
└── README.md
```

## Plug-and-play

- **Port**: `ports.Scheduler` — `Run(ctx, handler) error`; handler is `func(ctx) error` and runs on each tick/schedule.
- **Default adapter**: `internal/adapters/scheduler/ticker` — runs the job every N duration (e.g. 1 minute); use for local dev or tests.
- **Production**: implement `ports.Scheduler` in `pkg/scheduler/cron` (or cloud) and in `cmd/worker/main.go` set:
  ```go
  scheduler := cron.NewScheduler(cron.Config{Spec: "0 * * * *"})  // every hour
  ```
  No other code changes needed.

See `pkg/README.md` for the contract and an example cron adapter.

## How to use

1. Copy the template and set the `module` in `go.mod` (e.g. `github.com/your-org/your-worker`).
2. Replace or extend `internal/handler/handler.go` with your job logic (cleanup, sync, report, etc.).
3. For production, add a scheduler plugin (e.g. `pkg/scheduler/cron`) implementing `ports.Scheduler` and use it in `cmd/worker/main.go`.
4. Run:
   ```bash
   make run
   # or: go run ./cmd/worker
   ```
   With the default ticker, the worker runs the job once immediately and then every minute until SIGINT/SIGTERM.

## Build and tests

```bash
make build   # bin/worker
make run     # run worker
make test    # run tests
make lint    # golangci-lint (requires installation)
```

## License

As per the Cosmos Toolkit project.
