package repository

import (
	"context"
	"sync"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

// RepoMemoryImpl структура по работе с внутренней памятью
type RepoMemoryImpl struct {
	urlStorage map[string]string
	mu         sync.Mutex
}

// NewRepoMemoryImpl создание структуры по работе с внутренней памятью
func NewRepoMemoryImpl() *RepoMemoryImpl {
	return &RepoMemoryImpl{
		urlStorage: make(map[string]string),
	}
}

// CreateURL создание урла во внутренней памяти
func (r *RepoMemoryImpl) CreateURL(ctx context.Context, shortURL, originalURL, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, found := r.urlStorage[shortURL]; found {
		return custom_errs.ErrUniqueViolation
	}
	r.urlStorage[shortURL] = originalURL

	return nil
}

// GetURLByShortURL получение урла из внутренней памяти
func (r *RepoMemoryImpl) GetURLByShortURL(ctx context.Context, shortURL string) (*URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	shortURLs := make([]string, len(r.urlStorage))
	for url := range r.urlStorage {
		shortURLs = append(shortURLs, url)
	}

	for _, k := range shortURLs {
		if k == shortURL {
			findURL := &URL{
				ID:          int64(len(shortURLs)),
				ShortURL:    k,
				OriginalURL: r.urlStorage[k],
			}
			return findURL, nil
		}
	}

	return nil, custom_errs.ErrNotFound
}
