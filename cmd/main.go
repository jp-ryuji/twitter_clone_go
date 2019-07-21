package main

import (
	"net"
	"os"
	"os/signal"

	"github.com/jp-ryuji/twitter_clone_go/internal/di"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	logger := di.InjectLogger()

	port, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal("failed to open listener", zap.Error(err))
	}

	server := grpc.NewServer()

	closer := di.RegisterHandler(server)
	defer closer()

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
