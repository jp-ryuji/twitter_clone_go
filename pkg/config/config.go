package config

import (
	"io/ioutil"
	"strings"

	"golang.org/x/xerrors"
	yaml "gopkg.in/yaml.v2"
)

// AppEnv represents the application environment (e.g. dev, prod).
type AppEnv string

// Env variables represent the runtime environment.
const (
	EnvTest AppEnv = "test"
	EnvDev  AppEnv = "dev"
	EnvProd AppEnv = "prod"
)

// Config represents the configuration values.
type Config struct {
	Env      AppEnv
	Database Database `yaml:"database"`
}

// NewConfig returns a new Config for the given environment.
func NewConfig(env AppEnv) *Config {
	return &Config{Env: env}
}

// IsDev returns true if env is dev.
func (c *Config) IsDev() bool {
	return c.Env == EnvDev
}

// IsTest returns true if env is test.
func (c *Config) IsTest() bool {
	return c.Env == EnvTest
}

// IsProd returns true if env is prod.
func (c *Config) IsProd() bool {
	return c.Env == EnvProd
}

// Load loads a configuration from the specified path.
func (c *Config) Load(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return xerrors.Errorf("failed to read directory %s: %w", dirPath, err)
	}

	for _, file := range files {
		path := dirPath + "/" + file.Name()
		if file.IsDir() || strings.HasPrefix(file.Name(), ".") {
			continue
		}
		f, err := ioutil.ReadFile(path)
		if err != nil {
			return xerrors.Errorf("failed to read file %s: %w", file.Name(), err)
		}
		if err := yaml.Unmarshal(f, c); err != nil {
			return xerrors.Errorf("failed to load config file %s: %w", file.Name(), err)
		}
	}
	return nil
}
