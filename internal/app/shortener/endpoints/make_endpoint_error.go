package endpoints

import (
	"errors"
	"net/http"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

// makeEndpointError обработчик ошибок
func makeEndpointError(res http.ResponseWriter, err error) {
	var httpCode int
	switch {
	case errors.Is(err, custom_errs.ErrNotFound):
		httpCode = http.StatusNotFound
		res.WriteHeader(http.StatusNotFound)
	case errors.Is(err, custom_errs.ErrNoContent):
		httpCode = http.StatusNoContent
		res.WriteHeader(http.StatusNoContent)
	case errors.Is(err, custom_errs.ErrUniqueViolation):
		httpCode = http.StatusConflict
		res.WriteHeader(http.StatusConflict)
	case errors.Is(err, custom_errs.ErrValidate):
		httpCode = http.StatusBadRequest
		res.WriteHeader(http.StatusBadRequest)
	default:
		httpCode = http.StatusInternalServerError
		res.WriteHeader(http.StatusInternalServerError)
	}
	http.Error(res, err.Error(), httpCode)
}
