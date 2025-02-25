package service

import (
	"go.uber.org/zap"

	"github.com/sprint1/config"
)

type ServiceImpl struct {
	URLStorage map[string]string
	lg         *zap.SugaredLogger
	cfg        config.Config
}

func NewService(lg *zap.SugaredLogger, cfg config.Config) *ServiceImpl {
	return &ServiceImpl{URLStorage: make(map[string]string), lg: lg, cfg: cfg}
}
