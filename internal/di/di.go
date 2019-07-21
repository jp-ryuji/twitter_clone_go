package di

import (
	"fmt"
	"io"
	"log"
	"os"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type envVarName string

var (
	envAppEnv       envVarName = "APP_ENV"
	envAppConfigDir envVarName = "APP_CONFIG_DIR"
)

// RegisterHandler registers gRPC handlers.
func RegisterHandler(server *grpc.Server) (closer func()) {
	// TODO: Methods to register  gRPC handlers come here.

	return func() { closeAll() }
}

var closerFuncs = map[string]io.Closer{}

func registerCloser(key string, closer io.Closer) {
	if _, exists := closerFuncs[key]; exists {
		log.Fatalf("duplicate closer key: %s", key)
	}
	closerFuncs[key] = closer
}

func closeAll() {
	logger := InjectLogger().Named("app")
	for key, closer := range closerFuncs {
		if err := closer.Close(); err != nil {
			logger.Error(fmt.Sprintf("failed to close %s", key), zap.Error(err))
		}
	}
}

func fatalInitFailure(componentName string, err error) {
	log.Fatalf("failed to init %s: %+v", componentName, err)
}

func loadEnvVar(name envVarName) (value string) {
	value = os.Getenv(string(name))
	if value == "" {
		log.Fatalf("empty environment variable: %s", name)
	}
	return value
}
