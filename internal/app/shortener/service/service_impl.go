package service

import (
	"github.com/gorilla/securecookie"
	"go.uber.org/zap"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/repository"
	"github.com/sprint1/internal/app/shortener/workers"
)

type ServiceImpl struct {
	lg         *zap.SugaredLogger
	cfg        config.Config
	repo       repository.RepoBase
	cookie     *securecookie.SecureCookie
	workerPool *workers.WorkerPool
}

func NewService(lg *zap.SugaredLogger, cfg config.Config, repo repository.RepoBase, workerPool *workers.WorkerPool) *ServiceImpl {
	serviceImpl := &ServiceImpl{lg: lg, cfg: cfg, repo: repo, workerPool: workerPool}
	serviceImpl.cookie = newSecureCookie()
	return serviceImpl
}

func newSecureCookie() *securecookie.SecureCookie {
	var hashKey = []byte("very-very-very-very-secret-key32")
	var blockKey = []byte("a-lot-of-secret!")
	return securecookie.New(hashKey, blockKey)
}
