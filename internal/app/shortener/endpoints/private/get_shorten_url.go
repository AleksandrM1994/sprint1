package private

import (
	"context"
	"errors"
	"fmt"

	private "github.com/sprint1/internal/app/shortener/endpoints/private/proto"
	"github.com/sprint1/internal/app/shortener/endpoints/public/middleware"
)

func (c *Controller) GetShortenURL(ctx context.Context, req *private.GetShortenURLRequest) (*private.GetShortenURLResponse, error) {
	userIDValue := ctx.Value(middleware.UserID)
	_, ok := userIDValue.(string)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	shortURL, errSaveURL := c.service.SaveURL(ctx, req.OriginalUrl, "")
	if errSaveURL != nil {
		return nil, fmt.Errorf("save url: %w", errSaveURL)
	}

	return &private.GetShortenURLResponse{
		ShortUrl: shortURL,
	}, nil
}
