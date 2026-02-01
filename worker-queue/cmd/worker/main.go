package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/your-org/your-worker/internal/adapters/queue/memory"
	"github.com/your-org/your-worker/internal/app"
	"github.com/your-org/your-worker/internal/handler"
	"github.com/your-org/your-worker/internal/ports"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Plug-and-play: replace this consumer with your queue client from pkg/queue.
	// Example: consumer := rabbitmq.NewConsumer(rabbitmq.Config{URL: os.Getenv("RABBITMQ_URL")})
	var consumer ports.Consumer = memory.NewConsumer()

	msgHandler := handler.New()
	w := app.NewWorker(consumer, msgHandler.Handle)

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
