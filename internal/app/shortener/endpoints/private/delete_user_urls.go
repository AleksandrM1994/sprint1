package private

import (
	"context"
	"errors"
	"fmt"

	private "github.com/sprint1/internal/app/shortener/endpoints/private/proto"
	"github.com/sprint1/internal/app/shortener/endpoints/public/middleware"
)

func (c *Controller) DeleteUserURLs(ctx context.Context, req *private.DeleteUserURLsRequest) (*private.DeleteUserURLsResponse, error) {
	userIDValue := ctx.Value(middleware.UserID)
	userID, ok := userIDValue.(string)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	errDeleteUserURLs := c.service.DeleteUserURLs(ctx, userID, req.Urls)
	if errDeleteUserURLs != nil {
		return nil, fmt.Errorf("DeleteUserURLs error: %w", errDeleteUserURLs)
	}

	return nil, nil
}
