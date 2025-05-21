package endpoints

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sprint1/internal/app/shortener/endpoints/middleware"
	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

func (c *Controller) DeleteUserURLs(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
	defer cancel()

	userIDValue := ctx.Value(middleware.UserID)
	userID, ok := userIDValue.(string)
	if !ok {
		userID = ""
	}

	var urlsForDelete []string
	errDecode := json.NewDecoder(req.Body).Decode(&urlsForDelete)
	if errDecode != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errDecode.Error(), http.StatusInternalServerError)
		return
	}

	if len(urlsForDelete) == 0 || userID == "" {
		res.WriteHeader(http.StatusBadRequest)
		http.Error(res, custom_errs.ErrBadRequest.Error(), http.StatusBadRequest)
		return
	}

	errDeleteUserURLs := c.service.DeleteUserURLs(ctx, userID, urlsForDelete)
	if errDeleteUserURLs != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errDeleteUserURLs.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusAccepted)
}
