package endpoints

import (
	"context"
	"net/http"
	"time"
)

func (c *Controller) GetOriginalURLHandler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
	defer cancel()

	id := req.URL.Path[len("/"):]
	originalURL, err := c.service.GetOriginalURL(ctx, id)
	if err != nil {
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
