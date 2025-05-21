package errors

import "errors"

// ошибка уникальности записи
var ErrUniqueViolation = errors.New("unique violation error")

// ошибка отсутствия данных при поиске в БД
var ErrNotFound = errors.New("not found")

// ошибка отсутствия данных
var ErrNoContent = errors.New("not content")

// ошибка валидации
var ErrValidate = errors.New("empty value")

// ошибка авторизации
var ErrUnauthorized = errors.New("unauthorized")

// ошибка неверного запроса
var ErrBadRequest = errors.New("bad request")

// ошибка запрашиваемый ресурс больше не доступен
var ErrResourceGone = errors.New("resource gone")
