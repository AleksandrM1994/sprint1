package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/repository"
)

const (
	passwordLength = 20 // длина пароля
	fiveNums       = 5  // количество символов в пароле
)

// CreateUser сервисная функция по созданию пользователя
func (s *ServiceImpl) CreateUser(ctx context.Context, login string) (string, error) {
	if login == "" {
		return "", fmt.Errorf("failed validate login: %w", custom_errs.ErrValidate)
	}

	pass, err := password.Generate(passwordLength, fiveNums, fiveNums, false, false)
	if err != nil {
		return "", fmt.Errorf("failed to generate password: %w", err)
	}

	loginHash, err := s.HashData([]byte(login))
	if err != nil {
		return "", fmt.Errorf("failed to hash login: %w", err)
	}

	passwordHash, err := s.HashData([]byte(pass))
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	dbRepo, ok := s.repo.(repository.RepoDB)
	if !ok {
		return "", errors.New("failed to cast repo to repo.DB")
	}

	id := uuid.New().String()
	errCreateUser := dbRepo.CreateUser(ctx, id, loginHash, passwordHash)
	if errCreateUser != nil {
		return "", fmt.Errorf("failed to create user: %w", errCreateUser)
	}

	return pass, nil
}
