package private

import (
	"context"
	"errors"
	"fmt"

	private "github.com/sprint1/internal/app/shortener/endpoints/private/proto"
	"github.com/sprint1/internal/app/shortener/endpoints/public/middleware"
)

func (c *Controller) SaveURL(ctx context.Context, req *private.SaveURLRequest) (*private.SaveURLResponse, error) {
	userIDValue := ctx.Value(middleware.UserID)
	userID, ok := userIDValue.(string)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	shortURL, errSaveURL := c.service.SaveURL(ctx, req.Url, userID)
	if errSaveURL != nil {
		return nil, fmt.Errorf("save url: %w", errSaveURL)
	}

	return &private.SaveURLResponse{
		ShortUrl: shortURL,
	}, nil
}
