package private

import (
	"context"
	"errors"
	"fmt"

	private "github.com/sprint1/internal/app/shortener/endpoints/private/proto"
	"github.com/sprint1/internal/app/shortener/endpoints/public/middleware"
)

func (c *Controller) GetStatistics(ctx context.Context, _ *private.GetStatisticsRequest) (*private.GetStatisticsResponse, error) {
	userIDValue := ctx.Value(middleware.UserID)
	_, ok := userIDValue.(string)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	resGetStats, errGetStats := c.service.GetStats(ctx)
	if errGetStats != nil {
		return nil, fmt.Errorf("GetStats error: %w", errGetStats)
	}

	return &private.GetStatisticsResponse{
		UrlsCount:  resGetStats.URLs,
		UsersCount: resGetStats.Users,
	}, nil
}
