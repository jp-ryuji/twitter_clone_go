package config

import (
	"fmt"
	"net/url"
)

// Database represents a database configuration.
type Database struct {
	Driver           string            `yaml:"driver"`
	Host             string            `yaml:"host"`
	Port             string            `yaml:"port"`
	User             string            `yaml:"user"`
	Password         string            `yaml:"password"`
	Database         string            `yaml:"database"`
	Args             map[string]string `yaml:"args"`
	WaitForConnected int               `yaml:"waitForConnected"`
	MaxOpenConns     int               `yaml:"maxOpenConns"`
}

// ToDataSourceName returns a data source name.
func (db *Database) ToDataSourceName() string {
	args := url.Values{}
	for k, v := range db.Args {
		args.Set(k, v)
	}

	return fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s?%s",
		db.Driver,
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Database,
		args.Encode(),
	)
}
