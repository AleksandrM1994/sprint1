package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/repository"
)

func (s *ServiceImpl) AuthenticateUser(ctx context.Context, login, password string) (*repository.User, error) {
	if login == "" {
		return nil, fmt.Errorf("failed validate login: %w", custom_errs.ErrValidate)
	}

	if password == "" {
		return nil, fmt.Errorf("failed validate password: %w", custom_errs.ErrValidate)
	}

	loginHash, err := s.HashData([]byte(login))
	if err != nil {
		return nil, fmt.Errorf("failed to hash login: %w", err)
	}

	passwordHash, err := s.HashData([]byte(password))
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	dbRepo, ok := s.repo.(repository.RepoDB)
	if !ok {
		return nil, errors.New("failed to cast repo to repo.DB")
	}
	user, errGetUser := dbRepo.GetUser(ctx, loginHash, passwordHash)
	if errGetUser != nil {
		return nil, fmt.Errorf("failed get user: %w", errGetUser)
	}

	cookie, err := s.cookie.Encode(s.cfg.AuthUserCookieName, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to encode cookie: %w", err)
	}

	user.Cookie = cookie
	user.CookieFinish = DatePtr(time.Now().Add(24 * time.Hour))

	errUpdateUser := dbRepo.UpdateUser(ctx, user)
	if errUpdateUser != nil {
		return nil, fmt.Errorf("failed update user: %w", errUpdateUser)
	}

	return user, nil
}
