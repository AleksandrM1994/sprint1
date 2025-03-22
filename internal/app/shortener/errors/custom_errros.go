package errors

import "errors"

var ErrUniqueViolation = errors.New("unique violation error")
var ErrNotFound = errors.New("not found")
