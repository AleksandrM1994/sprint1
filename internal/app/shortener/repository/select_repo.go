package repository

import (
	"go.uber.org/zap"

	"github.com/sprint1/config"
)

func SelectRepo(lg *zap.SugaredLogger, cfg config.Config) (RepoBase, error) {
	if cfg.DNS != "" {
		db, err := Connect(cfg.DNS)
		if err != nil {
			return nil, err
		}
		errMigrate := Migrate(db)
		if errMigrate != nil {
			return nil, errMigrate
		}

		return &RepoDBImpl{lg: lg, cfg: cfg, db: db}, nil
	}

	return NewRepoMemoryImpl(), nil
}
