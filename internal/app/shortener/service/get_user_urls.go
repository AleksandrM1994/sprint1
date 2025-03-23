package service

import (
	"context"
	"errors"
	"fmt"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/repository"
)

type UserURLs struct {
	OriginalURL string
	ShortURL    string
}

func (s *ServiceImpl) GetUserURLs(ctx context.Context, userID string) ([]*UserURLs, error) {
	if userID == "" {
		return nil, custom_errs.ErrValidate
	}

	dbRepo, ok := s.repo.(repository.RepoDB)
	if !ok {
		return nil, errors.New("failed to cast repo to repo.DB")
	}

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
