package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/sprint1/config"
)

type RepoImpl struct {
	lg  *zap.SugaredLogger
	cfg config.Config
	db  *sqlx.DB
}

func NewRepoImpl(logger *zap.SugaredLogger, cfg config.Config) (*RepoImpl, error) {
	logger.Info(cfg.DNS)
	db, err := Connect(cfg.DNS)
	if err != nil {
		return nil, err
	}
	return &RepoImpl{
		lg:  logger,
		cfg: cfg,
		db:  db,
	}, nil
}

func Connect(dns string) (*sqlx.DB, error) {
	db, errConnect := sqlx.Connect("postgres", dns)
	if errConnect != nil {
		return nil, fmt.Errorf("sqlx.Connect: %w", errConnect)
	}
	return db, nil
}

func (r *RepoImpl) Ping() error {
	return r.db.Ping()
}
