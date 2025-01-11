package service

func getShortURL(url string) string {
	return url[:len(url)-1]
}
