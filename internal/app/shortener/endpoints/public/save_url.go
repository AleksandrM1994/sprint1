package public

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/sprint1/internal/app/shortener/endpoints/public/middleware"
	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

// SaveURLHandler ручка по сохранению урла
func (c *Controller) SaveURLHandler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
	defer cancel()

	request, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	userIDValue := ctx.Value(middleware.UserID)
	userID, ok := userIDValue.(string)
	if !ok {
		userID = ""
	}

	url := string(request)
	if url != "" {
		shortURL, errSaveURL := c.service.SaveURL(ctx, url, userID)
		switch {
		case errors.Is(errSaveURL, custom_errs.ErrUniqueViolation) && shortURL != "":
			res.WriteHeader(http.StatusConflict)
			_, _ = res.Write([]byte(c.cfg.BaseShortURL + "/" + shortURL))
		case errSaveURL != nil && shortURL == "":
			res.WriteHeader(http.StatusInternalServerError)
			http.Error(res, errSaveURL.Error(), http.StatusInternalServerError)
		case shortURL != "" && errSaveURL == nil:
			res.WriteHeader(http.StatusCreated)
			_, _ = res.Write([]byte(c.cfg.BaseShortURL + "/" + shortURL))
		case shortURL == "":
			res.WriteHeader(http.StatusBadRequest)
		default:
			res.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
