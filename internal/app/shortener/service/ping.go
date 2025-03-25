package service

import (
	"context"

	"github.com/sprint1/internal/app/shortener/repository"
)

func (s *ServiceImpl) Ping(ctx context.Context) error {
	if dbRepo, ok := s.repo.(repository.RepoDB); ok {
		return dbRepo.Ping(ctx)
	}
	return nil
}
