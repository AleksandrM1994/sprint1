package service

import "fmt"

func (s *ServiceImpl) GetOriginalURL(shortURL string) (string, error) {
	s.lg.Infow("GetOriginalURL request", "shortURL", shortURL)
	urlDB, err := s.repo.GetURLByShortURL(shortURL)
	if err != nil {
		return "", fmt.Errorf("repo.GetURLByShortURL: %v", err)
	}
	return urlDB.OriginalURL, nil
}
