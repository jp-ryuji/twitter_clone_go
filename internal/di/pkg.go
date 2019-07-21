package di

import (
	"github.com/jp-ryuji/twitter_clone_go/pkg/config"
	"github.com/jp-ryuji/twitter_clone_go/pkg/logging"

	"go.uber.org/zap"
)

var (
	_config *config.Config
)

// InjectConfig injects configuration based on the environment.
func InjectConfig() *config.Config {
	if _config != nil {
		return _config
	}

	env := loadEnvVar(envAppEnv)
	cfg := config.NewConfig(config.AppEnv(env))
	if err := cfg.Load(loadEnvVar(envAppConfigDir)); err != nil && !cfg.IsTest() {
		fatalInitFailure("app config", err)
	}
	_config = cfg
	return _config
}

// InjectLogger injects a logger.
func InjectLogger() *zap.Logger {
	cfg := InjectConfig()
	logger, err := logging.NewLogger(cfg.Env)
	if err != nil {
		fatalInitFailure("logger", err)
	}
	return logger
}
