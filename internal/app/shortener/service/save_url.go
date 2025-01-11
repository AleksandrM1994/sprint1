package service

func (s *ServiceImpl) SaveURL(url string) string {
	s.originalURL = url
	if len(url) > 0 {
		return getShortURL(url)
	}
	return ""
}
