package repository

import "context"

// RepoBase базовый интерфейс хранилища
type RepoBase interface {
	CreateURL(ctx context.Context, shortURL, originalURL, userID string) error
	GetURLByShortURL(ctx context.Context, shortURL string) (*URL, error)
}

// RepoDB интерфейс запросов в БД
type RepoDB interface {
	RepoBase
	Ping(ctx context.Context) error
	CreateURLs(ctx context.Context, urls []*URL) error
	CreateUser(ctx context.Context, id, login, password string) error
	GetUser(ctx context.Context, login, password string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id string) (*User, error)
	GetURLsByUserID(ctx context.Context, id string) ([]*URL, error)
	MakeURLsDeleted(ctx context.Context, urls []*URL) error
	GetStats(ctx context.Context) (uint32, uint32, error)
}
