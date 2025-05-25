package service

import (
	"context"

	"github.com/sprint1/internal/app/shortener/repository"
)

// Ping сервисная функция проверки жизни БД
func (s *ServiceImpl) Ping(ctx context.Context) error {
	if dbRepo, ok := s.repo.(repository.RepoDB); ok {
		return dbRepo.Ping(ctx)
	}
	return nil
}
