package private

import (
	"context"
	"errors"
	"fmt"

	private "github.com/sprint1/internal/app/shortener/endpoints/private/proto"
	"github.com/sprint1/internal/app/shortener/endpoints/public/middleware"
)

func (c *Controller) Ping(ctx context.Context, req *private.PingRequest) (*private.PingResponse, error) {
	userIDValue := ctx.Value(middleware.UserID)
	_, ok := userIDValue.(string)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	err := c.service.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return nil, nil
}
