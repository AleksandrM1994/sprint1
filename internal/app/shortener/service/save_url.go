package service

func (s *ServiceImpl) SaveURL(url string) string {
	if len(url) > 0 {
		shortUrl := getShortURL(url)
		s.OriginalURLsMap[shortUrl] = url
		return shortUrl
	}
	return ""
}
