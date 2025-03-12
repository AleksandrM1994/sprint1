package service

import (
	"go.uber.org/zap"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/repository"
)

type ServiceImpl struct {
	URLStorage map[string]string
	lg         *zap.SugaredLogger
	cfg        config.Config
	repo       repository.Repo
}

func NewService(lg *zap.SugaredLogger, cfg config.Config, repo repository.Repo) *ServiceImpl {
	return &ServiceImpl{URLStorage: make(map[string]string), lg: lg, cfg: cfg, repo: repo}
}
