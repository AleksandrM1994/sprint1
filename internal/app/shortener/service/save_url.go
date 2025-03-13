package service

import (
	"database/sql"
	"errors"
	"fmt"
)

func (s *ServiceImpl) SaveURL(url string) (string, error) {
	var shortURL string
	count := 0
	hashURL := HashString(url)
	fifthLength := len(hashURL) / 5

	for {
		// Обрезаем hashURL до нужной длины
		shortURL = hashURL[:fifthLength+count]

		urlDB, errGetURLByShortURL := s.repo.GetURLByShortURL(shortURL)
		if errGetURLByShortURL != nil && !errors.Is(errGetURLByShortURL, sql.ErrNoRows) {
			return "", fmt.Errorf("repo.GetURLByShortURL: %v", errGetURLByShortURL)
		}

		// Проверяем, существует ли уже этот короткий URL
		if urlDB == nil {
			// Если нет, сохраняем его и выходим из цикла
			errCreateURL := s.repo.CreateURL(shortURL, url)
			if errCreateURL != nil {
				return "", fmt.Errorf("repo.CreateURL:%w", errCreateURL)
			}
			break
		}

		// Увеличиваем count для следующей итерации
		count++

		// Проверяем, не достигли ли мы максимальной длины
		if fifthLength+count > len(hashURL) {
			// Если да, возвращаем пустую строку
			shortURL = ""
			break
		}
	}

	urlDB, errGetURLByShortURL := s.repo.GetURLByShortURL(shortURL)
	if errGetURLByShortURL != nil {
		return "", fmt.Errorf("repo.GetURLByShortURL: %v", errGetURLByShortURL)
	}

	errInsertURLInFile := s.InsertURLInFile(&URLInfo{
		UUID:        urlDB.Id,
		ShortURL:    urlDB.ShortURL,
		OriginalURL: urlDB.OriginalURL,
	})
	if errInsertURLInFile != nil {
		return "", fmt.Errorf("InsertURLInFile:%w", errInsertURLInFile)
	}

	return shortURL, nil
}
