package service

import "github.com/sprint1/internal/app/shortener/repository"

type Service interface {
	GetOriginalURL(url string) (string, error)
	SaveURL(userID, url string) (string, error)
	InsertURLInFile(URLInfo *URLInfo) error
	Ping() error
	CreateUser(login string) (string, error)
	AuthenticateUser(login, password string) (*repository.User, error)
	HashData(data []byte) (string, error)
	CheckCookie(cookie string) (string, error)
}
