package worker

import (
	"context"
	"log"

	"github.com/your-org/your-app/internal/ports"
)

// Worker runs the queue consumer with a message handler.
type Worker struct {
	consumer ports.Consumer
	handler  ports.MessageHandler
}

// New builds a worker with the given consumer and handler.
func New(consumer ports.Consumer, handler ports.MessageHandler) *Worker {
	return &Worker{consumer: consumer, handler: handler}
}

// Run starts consuming; blocks until ctx is cancelled.
func (w *Worker) Run(ctx context.Context) error {
	log.Println("worker: starting consumer")
	defer log.Println("worker: stopped")
	return w.consumer.Subscribe(ctx, w.handler)
}
