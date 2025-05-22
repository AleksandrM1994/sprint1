package workers

import (
	"context"

	"go.uber.org/zap"

	"github.com/sprint1/internal/app/shortener/repository"
)

// WorkerPool управляет воркерами и каналами
type WorkerPool struct {
	lg   *zap.SugaredLogger
	repo repository.RepoBase

	urls chan []*repository.URL
}

// NewWorkerPool создает новый пул воркеров
func NewWorkerPool(lg *zap.SugaredLogger, repo repository.RepoBase) *WorkerPool {
	return &WorkerPool{
		lg:   lg,
		urls: make(chan []*repository.URL, 10),
		repo: repo,
	}
}

// Start запускает воркеры
func (wp *WorkerPool) Start() {
	go wp.worker()
}

// worker функция, которая будет обрабатывать данные из канала
func (wp *WorkerPool) worker() {
	for urls := range wp.urls {
		dbRepo, _ := wp.repo.(repository.RepoDB)

		err := dbRepo.MakeURLsDeleted(context.Background(), urls)
		if err != nil {
			wp.lg.Infow("failed to make urls", "urls", urls, "err", err)
		}
	}
}

// Submit добавляет задачу в канал
func (wp *WorkerPool) Submit(url []*repository.URL) {
	// через отдельную горутину генератор отправляет данные в канал
	go func() {
		wp.urls <- url
	}()
}
