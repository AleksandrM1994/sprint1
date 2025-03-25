package endpoints

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

func (c *Controller) SaveURLHandler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
	defer cancel()

	request, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	url := string(request)
	if url != "" {
		shortURL, errSaveURL := c.service.SaveURL(ctx, url)
		switch {
		case errors.Is(errSaveURL, custom_errs.ErrUniqueViolation) && shortURL != "":
			res.WriteHeader(http.StatusConflict)
			_, writeErr := res.Write([]byte(c.cfg.BaseShortURL + "/" + shortURL))
			if writeErr != nil {
				res.WriteHeader(http.StatusInternalServerError)
				http.Error(res, writeErr.Error(), http.StatusInternalServerError)
				return
			}
		case errSaveURL != nil && shortURL == "":
			res.WriteHeader(http.StatusInternalServerError)
			http.Error(res, errSaveURL.Error(), http.StatusInternalServerError)
		case shortURL != "" && errSaveURL == nil:
			res.WriteHeader(http.StatusCreated)
			_, writeErr := res.Write([]byte(c.cfg.BaseShortURL + "/" + shortURL))
			if writeErr != nil {
				res.WriteHeader(http.StatusInternalServerError)
				http.Error(res, writeErr.Error(), http.StatusInternalServerError)
				return
			}
		case shortURL == "":
			res.WriteHeader(http.StatusBadRequest)
		default:
			res.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
