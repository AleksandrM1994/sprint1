package endpoints

import (
	"errors"
	"io"
	"net/http"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

func (c *Controller) SaveURLHandler(res http.ResponseWriter, req *http.Request) {
	request, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	url := string(request)
	if url != "" {
		shortUrl, errSaveURL := c.service.SaveURL(url)
		switch {
		case errors.Is(errSaveURL, custom_errs.ErrUniqueViolation) && shortUrl != "":
			res.WriteHeader(http.StatusConflict)
			res.Header().Set("Content-Type", "text/plain")
			_, writeErr := res.Write([]byte(c.cfg.BaseShortURL + "/" + shortUrl))
			if writeErr != nil {
				res.WriteHeader(http.StatusInternalServerError)
				http.Error(res, writeErr.Error(), http.StatusInternalServerError)
			}
		case errSaveURL != nil && shortUrl == "":
			res.WriteHeader(http.StatusInternalServerError)
			http.Error(res, errSaveURL.Error(), http.StatusInternalServerError)
		case shortUrl != "" && errSaveURL == nil:
			res.WriteHeader(http.StatusCreated)
			res.Header().Set("Content-Type", "text/plain")
			_, writeErr := res.Write([]byte(c.cfg.BaseShortURL + "/" + shortUrl))
			if writeErr != nil {
				res.WriteHeader(http.StatusInternalServerError)
				http.Error(res, writeErr.Error(), http.StatusInternalServerError)
			}
		case shortUrl == "":
			res.WriteHeader(http.StatusBadRequest)
		default:
			res.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
