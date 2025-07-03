package private

import (
	"context"
	"errors"
	"fmt"

	private "github.com/sprint1/internal/app/shortener/endpoints/private/proto"
	"github.com/sprint1/internal/app/shortener/endpoints/public/middleware"
)

func (c *Controller) GetOriginalURL(ctx context.Context, req *private.GetOriginalURLRequest) (*private.GetOriginalURLResponse, error) {
	userIDValue := ctx.Value(middleware.UserID)
	_, ok := userIDValue.(string)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	originalURL, err := c.service.GetOriginalURL(ctx, req.ShortId)
	if err != nil {
		return nil, fmt.Errorf("GetOriginalURL error: %w", err)
	}

	return &private.GetOriginalURLResponse{
		OriginalUrl: originalURL,
	}, nil
}
