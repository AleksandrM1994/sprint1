package service

import (
	"context"
	"errors"
	"fmt"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/repository"
)

// UserURLs структура с инфой об урле
type UserURLs struct {
	OriginalURL string
	ShortURL    string
}

// GetUserURLs сервисная функция по получению урлов пользователя
func (s *ServiceImpl) GetUserURLs(ctx context.Context, userID string) ([]*UserURLs, error) {
	dbRepo, ok := s.repo.(repository.RepoDB)
	if !ok {
		return nil, errors.New("failed to cast repo to repo.DB")
	}

	s.lg.Infow("GetUserURLsRequest", "userID", userID)

	urls, errGetURLsByUserID := dbRepo.GetURLsByUserID(ctx, userID)
	if errGetURLsByUserID != nil {
		return nil, fmt.Errorf("failed to fetch URLs by userID %q: %w", userID, errGetURLsByUserID)
	}

	if len(urls) == 0 {
		return nil, custom_errs.ErrNoContent
	}

	userURLs := make([]*UserURLs, 0)
	for _, url := range urls {
		userURLs = append(userURLs, &UserURLs{
			OriginalURL: url.OriginalURL,
			ShortURL:    s.cfg.BaseShortURL + "/" + url.ShortURL,
		})
	}

	return userURLs, nil
}
