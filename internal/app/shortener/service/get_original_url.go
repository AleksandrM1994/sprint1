package service

import (
	"context"
	"fmt"
)

func (s *ServiceImpl) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	s.lg.Infow("GetOriginalURL request", "shortURL", shortURL)
	urlDB, err := s.repo.GetURLByShortURL(ctx, shortURL)
	if err != nil {
		return "", fmt.Errorf("repo.GetURLByShortURL: %v", err)
	}
	return urlDB.OriginalURL, nil
}
