package service

type Service interface {
	GetOriginalURL(url string) (string, error)
	SaveURL(url string) (string, error)
	InsertURLInFile(URLInfo *URLInfo) error
	Ping() error
}
