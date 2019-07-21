package datastore

import (
	"database/sql"
	"time"

	"github.com/jp-ryuji/twitter_clone_go/pkg/config"
	"github.com/volatiletech/sqlboiler/boil"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

var (
	logger *zap.Logger
)

// NewDB returns a database from the specified config.
func NewDB(cfg *config.Config, l *zap.Logger) (*sql.DB, error) {
	logger = l
	dbCfg := cfg.Database

	dsn := dbCfg.ToDataSourceName()
	db, err := sql.Open(dbCfg.Driver, dsn)
	if err != nil {
		return nil, xerrors.Errorf("failed to open database %s: %w", dsn, err)
	}

	if dbCfg.WaitForConnected > 0 {
		if err = waitForConnected(db, dbCfg.WaitForConnected); err != nil {
			return nil, err
		}
	}

	db.SetMaxOpenConns(dbCfg.MaxOpenConns)
	db.SetMaxIdleConns(dbCfg.MaxOpenConns * 2)
	db.SetConnMaxLifetime(time.Duration(dbCfg.MaxOpenConns) * time.Second)

	boil.SetDB(db)
	if !cfg.IsProd() {
		boil.DebugMode = true
	}

	return db, nil
}

func waitForConnected(db *sql.DB, retry int) error {
	if err := db.Ping(); err == nil {
		return nil
	}
	retry--
	if retry < 0 {
		return xerrors.Errorf("failed to connect")
	}
	time.Sleep(1 * time.Second)
	return waitForConnected(db, retry)
}
