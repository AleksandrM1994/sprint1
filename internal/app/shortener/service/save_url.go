package service

import (
	"errors"
	"fmt"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/helpers"
)

func (s *ServiceImpl) SaveURL(url string) (string, error) {
	var shortURL string
	count := 0

	url = helpers.RemoveControlCharacters(url)
	hashURL := HashString(url)
	fifthLength := len(hashURL) / 5

	// Обрезаем hashURL до нужной длины
	shortURL = hashURL[:fifthLength+count]

	s.lg.Infow("SaveURL request", "url", url, "shortURL", shortURL)

	errCreateURL := s.repo.CreateURL(shortURL, url)
	if errCreateURL != nil {
		if errors.Is(errCreateURL, custom_errs.ErrUniqueViolation) {
			urlDB, errGetURLByShortURL := s.repo.GetURLByShortURL(shortURL)
			if errGetURLByShortURL != nil {
				return "", fmt.Errorf("repo.GetURLByShortURL: %v", errGetURLByShortURL)
			}
			return urlDB.ShortURL, errCreateURL
		}
		return "", fmt.Errorf("repo.CreateURL:%w", errCreateURL)
	}

	urlDB, errGetURLByShortURL := s.repo.GetURLByShortURL(shortURL)
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
