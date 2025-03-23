package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

const (
	passwordLength = 20
	fiveNums       = 5
)

func (s *ServiceImpl) CreateUser(login string) (string, error) {
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

	id := uuid.New().String()
	errCreateUser := s.repo.CreateUser(id, loginHash, passwordHash)
	if errCreateUser != nil {
		return "", fmt.Errorf("failed to create user: %w", errCreateUser)
	}

	return pass, nil
}
