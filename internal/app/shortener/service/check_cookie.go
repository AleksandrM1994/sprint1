package service

import (
	"fmt"
	"time"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

func (s *ServiceImpl) CheckCookie(cookie string) (string, error) {
	var userID string
	err := s.cookie.Decode(s.cfg.AuthUserCookieName, cookie, &userID)
	if err != nil {
		return "", fmt.Errorf("failed to decode cookie: %w", err)
	}

	if userID == "" {
		return "", custom_errs.ErrUnauthorized
	}

	user, errGetUserByID := s.repo.GetUserByID(userID)
	if errGetUserByID != nil {
		return "", fmt.Errorf("failed to get user by ID: %w", errGetUserByID)
	}

	if time.Now().After(*user.CookieFinish) || user.Cookie != cookie {
		return "", custom_errs.ErrUnauthorized
	}

	return userID, nil
}
