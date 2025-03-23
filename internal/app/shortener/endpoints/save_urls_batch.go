package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/sprint1/internal/app/shortener/endpoints/middleware"
)

type URLInBatch struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url,omitempty"`
	ShortURL      string `json:"short_url,omitempty"`
}

func (c *Controller) SaveURLsBatch(res http.ResponseWriter, req *http.Request) {
	var saveURLsBatchRequest []URLInBatch
	errDecode := json.NewDecoder(req.Body).Decode(&saveURLsBatchRequest)
	if errDecode != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errDecode.Error(), http.StatusInternalServerError)
		return
	}

	ctx := req.Context()
	userID := ctx.Value(middleware.UserID).(string)

	saveURLsBatchResponse := make([]URLInBatch, 0, len(saveURLsBatchRequest))
	if len(saveURLsBatchRequest) != 0 {
		for _, url := range saveURLsBatchRequest {
			shortURL, errSaveURL := c.service.SaveURL(userID, url.OriginalURL)
			if errSaveURL != nil {
				res.WriteHeader(http.StatusInternalServerError)
				http.Error(res, errSaveURL.Error(), http.StatusInternalServerError)
				return
			}
			if shortURL != "" {
				url.OriginalURL = ""
				url.ShortURL = c.cfg.BaseShortURL + "/" + shortURL
				saveURLsBatchResponse = append(saveURLsBatchResponse, url)
			} else {
				res.WriteHeader(http.StatusBadRequest)
			}
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		body, errMarshal := json.Marshal(saveURLsBatchResponse)
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
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
