package errors

import "errors"

var ErrUniqueViolation = errors.New("unique violation error")
var ErrNotFound = errors.New("not found")
var ErrNoContent = errors.New("not content")
var ErrValidate = errors.New("empty value")
var ErrUnauthorized = errors.New("unauthorized")
