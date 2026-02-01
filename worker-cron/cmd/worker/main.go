package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/your-org/your-worker/internal/adapters/scheduler/ticker"
	"github.com/your-org/your-worker/internal/app"
	"github.com/your-org/your-worker/internal/handler"
	"github.com/your-org/your-worker/internal/ports"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Plug-and-play: replace this scheduler with your cron/cloud client from pkg/scheduler.
	// Example: scheduler := cron.NewScheduler(cron.Config{Spec: "0 * * * *"})  // every hour
	var scheduler ports.Scheduler = ticker.NewScheduler(1 * time.Minute)

	jobHandler := handler.New()
	w := app.NewWorker(scheduler, jobHandler.Handle)

	go func() {
		if err := w.Run(ctx); err != nil && err != context.Canceled {
			log.Printf("worker error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down worker...")
	cancel()
}
