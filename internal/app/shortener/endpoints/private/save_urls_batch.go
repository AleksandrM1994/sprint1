package private

import (
	"context"
	"errors"
	"fmt"

	private "github.com/sprint1/internal/app/shortener/endpoints/private/proto"
	"github.com/sprint1/internal/app/shortener/endpoints/public/middleware"
	"github.com/sprint1/internal/app/shortener/service"
)

func (c *Controller) SaveURLsBatch(ctx context.Context, req *private.SaveURLsBatchRequest) (*private.SaveURLsBatchResponse, error) {
	userIDValue := ctx.Value(middleware.UserID)
	userID, ok := userIDValue.(string)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	urlsForSave := make([]*service.URLInBatch, 0)
	for _, url := range req.Urls {
		urlsForSave = append(urlsForSave, &service.URLInBatch{
			CorrelationID: url.CorrelationId,
			OriginalURL:   url.OriginalUrl,
		})
	}

	newURLs, errSaveURL := c.service.SaveURLsBatch(ctx, urlsForSave, userID)
	if errSaveURL != nil {
		return nil, fmt.Errorf("save urls: %w", errSaveURL)
	}

	savedURLs := make([]*private.URLBatch, 0)
	for _, url := range newURLs {
		savedURLs = append(savedURLs, &private.URLBatch{
			CorrelationId: url.CorrelationID,
			OriginalUrl:   url.OriginalURL,
			ShortUrl:      url.ShortURL,
		})
	}

	return &private.SaveURLsBatchResponse{
		Urls: savedURLs,
	}, nil
}
