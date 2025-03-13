package repository

type Repo interface {
	Ping() error
	CreateURL(shortURL string, originalURL string) error
	GetURLByShortURL(shortURL string) (*URL, error)
}
