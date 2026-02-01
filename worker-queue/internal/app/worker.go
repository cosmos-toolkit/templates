package app

import (
	"context"
	"log"

	"github.com/your-org/your-worker/internal/ports"
)

// Worker runs a queue consumer with a message handler (plug-and-play).
// Inject any Consumer implementation (memory, rabbitmq, sqs, kafka) via NewWorker.
type Worker struct {
	consumer ports.Consumer
	handler  ports.MessageHandler
}

// NewWorker builds a worker with the given consumer and handler.
// Plug your queue client: e.g. consumer := rabbitmq.NewConsumer(cfg) from pkg/queue/rabbitmq.
func NewWorker(consumer ports.Consumer, handler ports.MessageHandler) *Worker {
	return &Worker{consumer: consumer, handler: handler}
}

// Run starts consuming messages. It blocks until ctx is cancelled.
func (w *Worker) Run(ctx context.Context) error {
	log.Println("worker: starting consumer")
	defer log.Println("worker: stopped")
	return w.consumer.Subscribe(ctx, w.handler)
}
