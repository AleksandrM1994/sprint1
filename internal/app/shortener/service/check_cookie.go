package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/repository"
)

// CheckCookie сервисная функция по проверке куки
func (s *ServiceImpl) CheckCookie(ctx context.Context, cookie string) (string, error) {
	var userID string
	err := s.cookie.Decode(s.cfg.AuthUserCookieName, cookie, &userID)
	if err != nil {
		return "", fmt.Errorf("failed to decode cookie: %w", err)
	}

	if userID == "" {
		return "", custom_errs.ErrUnauthorized
	}

	dbRepo, ok := s.repo.(repository.RepoDB)
	if !ok {
		return "", errors.New("failed to cast repo to repo.DB")
	}

	user, errGetUserByID := dbRepo.GetUserByID(ctx, userID)
	if errGetUserByID != nil {
		return "", fmt.Errorf("failed to get user by ID: %w", errGetUserByID)
	}

	if time.Now().After(*user.CookieFinish) || user.Cookie != cookie {
		return "", custom_errs.ErrUnauthorized
	}

	return userID, nil
}
