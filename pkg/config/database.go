package config

import (
	"fmt"
	"net/url"
)

// Database represents a database configuration.
type Database struct {
	Driver           string
	Host             string
	Port             string
	User             string
	Password         string
	Database         string
	Args             map[string]string
	WaitForConnected int
	MaxOpenConns     int
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
