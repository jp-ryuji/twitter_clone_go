package di

import (
	"database/sql"
	"log"

	"github.com/jp-ryuji/twitter_clone_go/internal/infrastructure/datastore"
)

var (
	_db *sql.DB
)

// InjectDB returns singleton database instance. Being singleton is guaranteed by *sql.DB.
func InjectDB() *sql.DB {
	if _db != nil {
		return _db
	}

	logger := InjectLogger().Named("repository")
	db, err := datastore.NewDB(InjectConfig(), logger)
	if err != nil {
		log.Fatalf("failed to init database: %+v", err)
	}
	registerCloser("db", db)
	_db = db
	return _db
}
