package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/repository"
)

func (s *ServiceImpl) DeleteUserURLs(ctx context.Context, userID string, urls []string) (chan error, error) {
	dbRepo, ok := s.repo.(repository.RepoDB)
	if !ok {
		return nil, errors.New("failed to cast repo to repo.DB")
	}

	urlsForDelete := make([]*repository.URL, 0, len(urls))
	for _, shortURL := range urls {
		urlDB, errGetURLByShortURL := dbRepo.GetURLByShortURL(ctx, shortURL)
		if errGetURLByShortURL != nil {
			return nil, fmt.Errorf("err get url by short url: %w", errGetURLByShortURL)
		}

		if urlDB.UserID != userID {
			return nil, custom_errs.ErrBadRequest
		}

		urlDB.IsDeleted = true
		urlsForDelete = append(urlsForDelete, urlDB)
	}

	errChan := make(chan error)
	var wg sync.WaitGroup

	for _, url := range urlsForDelete {
		wg.Add(1)
		go func(url *repository.URL) {
			defer wg.Done()

			errMakeURLsDeleted := dbRepo.MakeURLsDeleted(ctx, []*repository.URL{url})
			errChan <- errMakeURLsDeleted
		}(url)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	return errChan, nil
}
