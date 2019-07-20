package main

import (
	"net"
	"os"
	"os/signal"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	// FIXME: Inject it based on the development mode.
	logger, _ := zap.NewDevelopment()

	port, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal("failed to open listener", zap.Error(err))
	}

	server := grpc.NewServer()

	go func() {
		logger.Info("start gRPC server")
		if err := server.Serve(port); err != nil {
			logger.Fatal("failed to serve: %+v", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("shouting down the server gracefully")
	server.GracefulStop()
}
