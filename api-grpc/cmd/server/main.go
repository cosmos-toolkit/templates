package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	entityv1 "github.com/your-org/your-app/api/pb/entity/v1"
	"github.com/your-org/your-app/internal/adapters/grpcserver"
	"github.com/your-org/your-app/internal/app"
)

func main() {
	addr := ":9090"
	if v := os.Getenv("GRPC_ADDR"); v != "" {
		addr = v
	}

	application := app.New()
	adapter := grpcserver.NewServer(application)

	gs := grpc.NewServer()
	entityv1.RegisterEntityServiceServer(gs, adapter)
	reflection.Register(gs)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	go func() {
		log.Printf("gRPC server listening on %s", addr)
		if err := gs.Serve(lis); err != nil {
			log.Fatalf("serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	gs.GracefulStop()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	application.Shutdown(ctx)
	log.Println("server stopped")
}
