package service

import (
	"context"
	"fmt"

	"github.com/sprint1/internal/app/shortener/repository"
)

// URLInBatch сервисная структура для ответа функции SaveURLsBatch
type URLInBatch struct {
	CorrelationID string
	OriginalURL   string
	ShortURL      string
}

// SaveURLsBatch сервисная функция по сохранению урлов в рамках одной транзакции
func (s *ServiceImpl) SaveURLsBatch(ctx context.Context, urls []*URLInBatch, userID string) ([]*URLInBatch, error) {
	var newURLs []*URLInBatch
	if dbRepo, ok := s.repo.(repository.RepoDB); ok {
		var urlsDB []*repository.URL
		for _, url := range urls {
			shortURL := CreateShortURL(url.OriginalURL)

			urlsDB = append(urlsDB, &repository.URL{
				ShortURL:    shortURL,
				OriginalURL: url.OriginalURL,
				UserID:      userID,
			})

			newURLs = append(newURLs, &URLInBatch{
				CorrelationID: url.CorrelationID,
				ShortURL:      s.cfg.BaseShortURL + "/" + shortURL,
			})
		}
		err := dbRepo.CreateURLs(ctx, urlsDB)
		if err != nil {
			return nil, fmt.Errorf("save urls batch: %w", err)
		}

		for _, url := range urlsDB {
			errInsertURLInFile := s.InsertURLInFile(&URLInfo{
				UUID:        url.ID,
				ShortURL:    url.ShortURL,
				OriginalURL: url.OriginalURL,
			})
			if errInsertURLInFile != nil {
				return nil, fmt.Errorf("InsertURLInFile:%w", errInsertURLInFile)
			}
		}
		return newURLs, nil
	}
	return nil, nil
}
