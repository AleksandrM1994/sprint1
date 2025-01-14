package service

func (s *ServiceImpl) GetOriginalURL(shortURL string) string {
	if originalURL, ok := s.OriginalURLsMap[shortURL]; ok {
		return originalURL
	}
	return ""
}
