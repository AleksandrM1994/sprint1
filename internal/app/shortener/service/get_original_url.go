package service

func (s *ServiceImpl) GetOriginalURL(shortUrl string) string {
	if originalURL, ok := s.URLStorage[shortUrl]; ok {
		return originalURL
	}
	return ""
}
