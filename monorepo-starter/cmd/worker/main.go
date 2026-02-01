package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/your-org/your-app/internal/adapters/queue/memory"
	"github.com/your-org/your-app/internal/ports"
	"github.com/your-org/your-app/internal/worker"
	"github.com/your-org/your-app/internal/worker/handler"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Plug-and-play: replace with pkg/queue/rabbitmq etc. in production.
	var consumer ports.Consumer = memory.NewConsumer()

	h := handler.New()
	w := worker.New(consumer, h.Handle)

	go func() {
		if err := w.Run(ctx); err != nil && err != context.Canceled {
			log.Printf("worker: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("worker: shutting down...")
	cancel()
}
