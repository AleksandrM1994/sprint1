package repository

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/sprint1/config"
)

func SelectRepo(lg *zap.SugaredLogger, cfg config.Config) (Repo, error) {
	if cfg.DNS != "" {
		repo, errNewRepoImpl := NewRepoDBImpl(lg, cfg)
		if errNewRepoImpl != nil {
			return nil, fmt.Errorf("repository.NewRepoImpl:%w", errNewRepoImpl)
		}
		return repo, nil
	}

	return NewRepoMemoryImpl(), nil
}
