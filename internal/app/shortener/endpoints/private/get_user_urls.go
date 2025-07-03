package private

import (
	"context"
	"errors"
	"fmt"

	private "github.com/sprint1/internal/app/shortener/endpoints/private/proto"
	"github.com/sprint1/internal/app/shortener/endpoints/public/middleware"
)

func (c *Controller) GetUserURLs(ctx context.Context, req *private.GetUserURLsRequest) (*private.GetUserURLsResponse, error) {
	userIDValue := ctx.Value(middleware.UserID)
	userID, ok := userIDValue.(string)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	urls, errGetUserURLs := c.service.GetUserURLs(ctx, userID)
	if errGetUserURLs != nil {
		return nil, fmt.Errorf("GetUserURLs error: %w", errGetUserURLs)
	}

	userURLs := make([]*private.UserURL, 0, len(urls))
	for _, url := range urls {
		userURLs = append(userURLs, &private.UserURL{
			OriginalUrl: url.OriginalURL,
			ShortUrl:    url.ShortURL,
		})
	}

	return &private.GetUserURLsResponse{
		UserUrls: userURLs,
	}, nil
}
