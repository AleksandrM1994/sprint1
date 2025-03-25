package repository

import "context"

type RepoBase interface {
	CreateURL(ctx context.Context, shortURL string, originalURL string) error
	GetURLByShortURL(ctx context.Context, shortURL string) (*URL, error)
}

type RepoDB interface {
	RepoBase
	Ping(ctx context.Context) error
	CreateURLs(ctx context.Context, urls []*URL) error
}
