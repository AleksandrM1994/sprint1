package service

import (
	"context"
	"errors"
	"fmt"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

func (s *ServiceImpl) SaveURL(ctx context.Context, url string) (string, error) {
	shortURL := CreateShortURL(url)

	s.lg.Infow("SaveURL request", "url", url, "shortURL", shortURL)

	errCreateURL := s.repo.CreateURL(ctx, shortURL, url)
	if errCreateURL != nil {
		if errors.Is(errCreateURL, custom_errs.ErrUniqueViolation) {
			urlDB, errGetURLByShortURL := s.repo.GetURLByShortURL(ctx, shortURL)
			if errGetURLByShortURL != nil {
				return "", fmt.Errorf("repo.GetURLByShortURL: %v", errGetURLByShortURL)
			}
			return urlDB.ShortURL, errCreateURL
		}
		return "", fmt.Errorf("repo.CreateURL:%w", errCreateURL)
	}

	urlDB, errGetURLByShortURL := s.repo.GetURLByShortURL(ctx, shortURL)
	if errGetURLByShortURL != nil {
		return "", fmt.Errorf("repo.GetURLByShortURL: %v", errGetURLByShortURL)
	}

	errInsertURLInFile := s.InsertURLInFile(&URLInfo{
		UUID:        urlDB.ID,
		ShortURL:    urlDB.ShortURL,
		OriginalURL: urlDB.OriginalURL,
	})
	if errInsertURLInFile != nil {
		return "", fmt.Errorf("InsertURLInFile:%w", errInsertURLInFile)
	}

	return shortURL, nil
}
