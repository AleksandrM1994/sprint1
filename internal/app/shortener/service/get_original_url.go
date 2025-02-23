package service

func (s *ServiceImpl) GetOriginalURL(shortUrl string) string {
	if originalURL, ok := s.OriginalURLsMap[shortUrl]; ok {
		return originalURL
	}
	return ""
}
