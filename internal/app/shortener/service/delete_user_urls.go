package service

import (
	"context"
	"errors"
	"fmt"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/repository"
)

func (s *ServiceImpl) DeleteUserURLs(ctx context.Context, userID string, urls []string) error {
	dbRepo, ok := s.repo.(repository.RepoDB)
	if !ok {
		return errors.New("failed to cast repo to repo.DB")
	}

	urlsForDelete := make([]*repository.URL, 0, len(urls))
	for _, shortURL := range urls {
		urlDB, errGetURLByShortURL := dbRepo.GetURLByShortURL(ctx, shortURL)
		if errGetURLByShortURL != nil {
			return fmt.Errorf("err get url by short url: %w", errGetURLByShortURL)
		}

		if urlDB.UserID != userID {
			return custom_errs.ErrBadRequest
		}

		urlDB.IsDeleted = true
		urlsForDelete = append(urlsForDelete, urlDB)
	}

	s.workerPool.Submit(urlsForDelete)

	return nil
}
