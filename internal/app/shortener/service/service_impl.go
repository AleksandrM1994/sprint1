package service

import "go.uber.org/zap"

type ServiceImpl struct {
	URLStorage map[string]string
	lg         *zap.SugaredLogger
}

func NewService(lg *zap.SugaredLogger) *ServiceImpl {
	return &ServiceImpl{URLStorage: make(map[string]string), lg: lg}
}
