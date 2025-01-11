package service

func (s *ServiceImpl) GetOriginalURL(url string) string {
	if getShortURL(s.originalURL) == url {
		return s.originalURL
	}
	return ""
}
