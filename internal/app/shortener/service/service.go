package service

type Service interface {
	GetOriginalURL(url string) string
	SaveURL(url string) string
	InsertURLInFile(URLInfo *URLInfo) error
}
