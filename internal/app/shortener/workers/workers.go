package workers

import (
	"context"

	"github.com/sprint1/internal/app/shortener/repository"
)

// WorkerPool управляет воркерами и каналами
type WorkerPool struct {
	jobs       chan []*repository.URL
	results    chan error
	repo       repository.RepoBase
	numWorkers int
}

// NewWorkerPool создает новый пул воркеров
func NewWorkerPool(numWorkers int, repo repository.RepoBase) *WorkerPool {
	return &WorkerPool{
		jobs:       make(chan []*repository.URL, 10),
		results:    make(chan error, 10),
		repo:       repo,
		numWorkers: numWorkers,
	}
}

// Start запускает воркеры
func (wp *WorkerPool) Start() {
	for w := 1; w <= wp.numWorkers; w++ {
		go wp.worker(w)
	}
}

// worker функция, которая будет обрабатывать данные из канала
func (wp *WorkerPool) worker(id int) {
	for urls := range wp.jobs {
		dbRepo, _ := wp.repo.(repository.RepoDB)

		err := dbRepo.MakeURLsDeleted(context.Background(), urls)
		wp.results <- err
	}
}

// Submit добавляет задачу в канал
func (wp *WorkerPool) Submit(url []*repository.URL) {
	wp.jobs <- url
}

// Close закрывает каналы
func (wp *WorkerPool) Close() {
	close(wp.jobs)
}
