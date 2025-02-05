package service

import "go.uber.org/zap"

type ServiceImpl struct {
	OriginalURLsMap map[string]string
	lg              *zap.SugaredLogger
}

func NewService(lg *zap.SugaredLogger) *ServiceImpl {
	return &ServiceImpl{OriginalURLsMap: make(map[string]string), lg: lg}
}
