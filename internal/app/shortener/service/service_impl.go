package service

import (
	"go.uber.org/zap"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/repository"
)

type ServiceImpl struct {
	lg   *zap.SugaredLogger
	cfg  config.Config
	repo repository.RepoBase
}

func NewService(lg *zap.SugaredLogger, cfg config.Config, repo repository.RepoBase) *ServiceImpl {
	return &ServiceImpl{lg: lg, cfg: cfg, repo: repo}
}
