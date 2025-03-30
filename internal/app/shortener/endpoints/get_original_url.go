package endpoints

import (
	"context"
	"errors"
	"net/http"
	"time"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

func (c *Controller) GetOriginalURLHandler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
	defer cancel()

	id := req.URL.Path[len("/"):]
	originalURL, err := c.service.GetOriginalURL(ctx, id)
	if err != nil {
		if errors.Is(err, custom_errs.ErrResourceGone) {
			res.WriteHeader(http.StatusGone)
			return
		}
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	if originalURL != "" {
		res.Header().Set("Location", originalURL)
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
