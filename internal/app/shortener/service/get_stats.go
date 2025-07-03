package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sprint1/internal/app/shortener/repository"
)

type GetStatsResponse struct {
	URLs  uint32
	Users uint32
}

// GetStats сервисная функция по получению статистики
func (s *ServiceImpl) GetStats(ctx context.Context) (*GetStatsResponse, error) {
	dbRepo, ok := s.repo.(repository.RepoDB)
	if !ok {
		return nil, errors.New("failed to cast repo to repo.DB")
	}

	usersCount, urlsCount, err := dbRepo.GetStats(ctx)
	if err != nil {
		return nil, fmt.Errorf("dbRepo.GetStats: %w", err)
	}

	return &GetStatsResponse{
		URLs:  urlsCount,
		Users: usersCount,
	}, nil
}
