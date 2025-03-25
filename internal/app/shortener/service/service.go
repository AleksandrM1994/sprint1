package service

import (
	"context"
)

type Service interface {
	GetOriginalURL(ctx context.Context, url string) (string, error)
	SaveURL(ctx context.Context, url string) (string, error)
	InsertURLInFile(URLInfo *URLInfo) error
	Ping(ctx context.Context) error
	SaveURLsBatch(ctx context.Context, urls []*URLInBatch) ([]*URLInBatch, error)
}
