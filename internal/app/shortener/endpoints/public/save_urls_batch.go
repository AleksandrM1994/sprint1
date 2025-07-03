package public

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/jinzhu/copier"

	"github.com/sprint1/internal/app/shortener/endpoints/public/middleware"
	"github.com/sprint1/internal/app/shortener/service"
)

// URLInBatch структура для урлов, которых будут сохранены в рамках одного батча
type URLInBatch struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url,omitempty"`
	ShortURL      string `json:"short_url,omitempty"`
}

// SaveURLsBatch ручка по сохранению урлов в рамках одного батча
func (c *Controller) SaveURLsBatch(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
	defer cancel()

	var saveURLsBatchRequest []*URLInBatch
	errDecode := json.NewDecoder(req.Body).Decode(&saveURLsBatchRequest)
	if errDecode != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errDecode.Error(), http.StatusInternalServerError)
		return
	}

	if len(saveURLsBatchRequest) == 0 {
		res.WriteHeader(http.StatusBadRequest)
		http.Error(res, errors.New("empty request").Error(), http.StatusBadRequest)
		return
	}

	var urls []*service.URLInBatch
	errCopy := copier.Copy(&urls, saveURLsBatchRequest)
	if errCopy != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errCopy.Error(), http.StatusInternalServerError)
		return
	}

	userIDValue := ctx.Value(middleware.UserID)
	userID, ok := userIDValue.(string)
	if !ok {
		userID = ""
	}

	newURLs, errSaveURL := c.service.SaveURLsBatch(ctx, urls, userID)
	if errSaveURL != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errSaveURL.Error(), http.StatusInternalServerError)
		return
	}

	var urlsForRes []URLInBatch
	errCopy = copier.Copy(&urlsForRes, newURLs)
	if errCopy != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errCopy.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	body, errMarshal := json.Marshal(urlsForRes)
	if errMarshal != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errMarshal.Error(), http.StatusInternalServerError)
		return
	}
	_, errWrite := res.Write(body)
	if errWrite != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errWrite.Error(), http.StatusInternalServerError)
		return
	}
}
