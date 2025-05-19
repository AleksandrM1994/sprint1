package service

import (
	"context"
	"fmt"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

func (s *ServiceImpl) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	s.lg.Infow("GetOriginalURL request", "shortURL", shortURL)
	urlDB, err := s.repo.GetURLByShortURL(ctx, shortURL)
	if err != nil {
		return "", fmt.Errorf("repo.GetURLByShortURL: %v", err)
	}
	if urlDB.IsDeleted {
		return "", custom_errs.ErrResourceGone
	}
	return urlDB.OriginalURL, nil
}
