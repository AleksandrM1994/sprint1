package service

import "encoding/base64"

func (s *ServiceImpl) SaveURL(url string) string {
	urlInBase64 := base64.URLEncoding.EncodeToString([]byte(url))
	s.OriginalURLsMap[urlInBase64] = url
	if urlInBase64 != "" {
		return urlInBase64
	}
	return ""
}
