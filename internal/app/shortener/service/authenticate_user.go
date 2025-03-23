package service

import (
	"fmt"
	"time"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/repository"
)

func (s *ServiceImpl) AuthenticateUser(login, password string) (*repository.User, error) {
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

	user, errGetUser := s.repo.GetUser(loginHash, passwordHash)
	if errGetUser != nil {
		return nil, fmt.Errorf("failed get user: %w", errGetUser)
	}

	cookie, err := s.cookie.Encode(s.cfg.AuthUserCookieName, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to encode cookie: %w", err)
	}

	user.Cookie = cookie
	user.CookieFinish = DatePtr(time.Now().Add(24 * time.Hour))

	errUpdateUser := s.repo.UpdateUser(user)
	if errUpdateUser != nil {
		return nil, fmt.Errorf("failed update user: %w", errUpdateUser)
	}

	return user, nil
}
