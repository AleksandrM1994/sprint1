package service

func (s *ServiceImpl) SaveURL(url string) string {
	if len(url) > 0 {
		shortURL := getShortURL(url)
		s.OriginalURLsMap[shortURL] = url
		return shortURL
	}
	return ""
}
