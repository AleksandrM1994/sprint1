package service

import (
	"context"

	"github.com/sprint1/internal/app/shortener/repository"
)

type Service interface {
	GetOriginalURL(ctx context.Context, url string) (string, error)
	SaveURL(ctx context.Context, url, userID string) (string, error)
	InsertURLInFile(URLInfo *URLInfo) error
	Ping(ctx context.Context) error
	SaveURLsBatch(ctx context.Context, urls []*URLInBatch, userID string) ([]*URLInBatch, error)
	CreateUser(ctx context.Context, login string) (string, error)
	AuthenticateUser(ctx context.Context, login, password string) (*repository.User, error)
	CheckCookie(ctx context.Context, cookie string) (string, error)
	GetUserURLs(ctx context.Context, userID string) ([]*UserURLs, error)

	HashData(data []byte) (string, error)
}
